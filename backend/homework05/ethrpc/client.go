package ethrpc

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const SepoliaChainID int64 = 11155111

var weiPerETH = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

type Client struct {
	rpc *ethclient.Client
}

type BlockInfo struct {
	Number       *big.Int
	Hash         common.Hash
	ParentHash   common.Hash
	Timestamp    uint64
	Transactions int
}

type TransactionInfo struct {
	Hash      common.Hash
	From      common.Address
	To        *common.Address
	Value     *big.Int
	Nonce     uint64
	Gas       uint64
	Type      uint8
	GasPrice  *big.Int
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Pending   bool
}

func Dial(ctx context.Context, rpcURL string) (*Client, error) {
	rpcURL = strings.TrimSpace(rpcURL)
	if rpcURL == "" {
		return nil, errors.New("RPC URL is required")
	}
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("dial ethereum RPC: %w", err)
	}
	return &Client{rpc: client}, nil
}

func (c *Client) Close() {
	if c != nil && c.rpc != nil {
		c.rpc.Close()
	}
}

func (c *Client) RPC() *ethclient.Client {
	return c.rpc
}

func (c *Client) Block(ctx context.Context, number *big.Int) (*BlockInfo, error) {
	block, err := c.rpc.BlockByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("get block: %w", err)
	}
	return &BlockInfo{
		Number:       new(big.Int).Set(block.Number()),
		Hash:         block.Hash(),
		ParentHash:   block.ParentHash(),
		Timestamp:    block.Time(),
		Transactions: len(block.Transactions()),
	}, nil
}

func (c *Client) Transaction(ctx context.Context, hash common.Hash) (*TransactionInfo, error) {
	tx, pending, err := c.rpc.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("get transaction: %w", err)
	}

	chainID := tx.ChainId()
	from, err := types.Sender(types.LatestSignerForChainID(chainID), tx)
	if err != nil {
		return nil, fmt.Errorf("recover transaction sender: %w", err)
	}

	return &TransactionInfo{
		Hash:      tx.Hash(),
		From:      from,
		To:        tx.To(),
		Value:     new(big.Int).Set(tx.Value()),
		Nonce:     tx.Nonce(),
		Gas:       tx.Gas(),
		Type:      tx.Type(),
		GasPrice:  new(big.Int).Set(tx.GasPrice()),
		GasTipCap: new(big.Int).Set(tx.GasTipCap()),
		GasFeeCap: new(big.Int).Set(tx.GasFeeCap()),
		Pending:   pending,
	}, nil
}

func (c *Client) SendETH(ctx context.Context, privateKeyHex string, to common.Address, amountWei *big.Int) (common.Hash, common.Address, error) {
	if amountWei == nil || amountWei.Sign() <= 0 {
		return common.Hash{}, common.Address{}, errors.New("transfer amount must be positive")
	}

	privateKey, err := PrivateKeyFromHex(privateKeyHex)
	if err != nil {
		return common.Hash{}, common.Address{}, err
	}
	from := crypto.PubkeyToAddress(privateKey.PublicKey)

	chainID, err := c.rpc.ChainID(ctx)
	if err != nil {
		return common.Hash{}, common.Address{}, fmt.Errorf("get chain ID: %w", err)
	}
	if err := requireSepolia(chainID); err != nil {
		return common.Hash{}, common.Address{}, err
	}
	nonce, err := c.rpc.PendingNonceAt(ctx, from)
	if err != nil {
		return common.Hash{}, common.Address{}, fmt.Errorf("get pending nonce: %w", err)
	}
	gasTipCap, err := c.rpc.SuggestGasTipCap(ctx)
	if err != nil {
		return common.Hash{}, common.Address{}, fmt.Errorf("suggest gas tip cap: %w", err)
	}
	header, err := c.rpc.HeaderByNumber(ctx, nil)
	if err != nil {
		return common.Hash{}, common.Address{}, fmt.Errorf("get latest header: %w", err)
	}
	if header.BaseFee == nil {
		return common.Hash{}, common.Address{}, errors.New("latest block has no EIP-1559 base fee")
	}
	gasFeeCap := new(big.Int).Add(
		new(big.Int).Mul(header.BaseFee, big.NewInt(2)),
		gasTipCap,
	)

	tx := newDynamicFeeTransfer(chainID, nonce, to, amountWei, gasTipCap, gasFeeCap)
	signed, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return common.Hash{}, common.Address{}, fmt.Errorf("sign transfer: %w", err)
	}
	if err := c.rpc.SendTransaction(ctx, signed); err != nil {
		return common.Hash{}, common.Address{}, fmt.Errorf("send transfer: %w", err)
	}
	return signed.Hash(), from, nil
}

func newDynamicFeeTransfer(chainID *big.Int, nonce uint64, to common.Address, amountWei, gasTipCap, gasFeeCap *big.Int) *types.Transaction {
	toCopy := to
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   new(big.Int).Set(chainID),
		Nonce:     nonce,
		GasTipCap: new(big.Int).Set(gasTipCap),
		GasFeeCap: new(big.Int).Set(gasFeeCap),
		Gas:       21_000,
		To:        &toCopy,
		Value:     new(big.Int).Set(amountWei),
	})
}

func (c *Client) NewTransactor(ctx context.Context, privateKeyHex string) (*bind.TransactOpts, common.Address, error) {
	privateKey, err := PrivateKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, common.Address{}, err
	}
	chainID, err := c.rpc.ChainID(ctx)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("get chain ID: %w", err)
	}
	if err := requireSepolia(chainID); err != nil {
		return nil, common.Address{}, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("create transactor: %w", err)
	}
	auth.Context = ctx
	return auth, crypto.PubkeyToAddress(privateKey.PublicKey), nil
}

func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c.rpc, tx)
	if err != nil {
		return nil, fmt.Errorf("wait for transaction %s: %w", tx.Hash().Hex(), err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return receipt, fmt.Errorf("transaction %s reverted", tx.Hash().Hex())
	}
	return receipt, nil
}

func requireSepolia(chainID *big.Int) error {
	if chainID == nil {
		return errors.New("chain ID is nil")
	}
	want := big.NewInt(SepoliaChainID)
	if chainID.Cmp(want) != 0 {
		return fmt.Errorf("write operations require Sepolia chain ID %d, got %s", SepoliaChainID, chainID)
	}
	return nil
}

func PrivateKeyFromHex(value string) (*ecdsa.PrivateKey, error) {
	value = strings.TrimSpace(value)
	value = strings.TrimPrefix(value, "0x")
	if value == "" {
		return nil, errors.New("private key is required")
	}
	key, err := crypto.HexToECDSA(value)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}
	return key, nil
}

func ParseETH(value string) (*big.Int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil, errors.New("ETH amount is required")
	}
	if strings.HasPrefix(value, "+") {
		value = strings.TrimPrefix(value, "+")
	}
	if strings.HasPrefix(value, "-") {
		return nil, errors.New("ETH amount must be positive")
	}

	parts := strings.Split(value, ".")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid ETH amount %q", value)
	}
	whole := parts[0]
	if whole == "" {
		whole = "0"
	}
	if !digitsOnly(whole) {
		return nil, fmt.Errorf("invalid ETH amount %q", value)
	}

	fraction := ""
	if len(parts) == 2 {
		fraction = parts[1]
		if fraction == "" || !digitsOnly(fraction) {
			return nil, fmt.Errorf("invalid ETH amount %q", value)
		}
		if len(fraction) > 18 {
			return nil, errors.New("ETH amount supports at most 18 decimal places")
		}
	}

	wholeWei, ok := new(big.Int).SetString(whole, 10)
	if !ok {
		return nil, fmt.Errorf("invalid ETH amount %q", value)
	}
	wholeWei.Mul(wholeWei, weiPerETH)

	if fraction != "" {
		fraction += strings.Repeat("0", 18-len(fraction))
		fractionWei, ok := new(big.Int).SetString(fraction, 10)
		if !ok {
			return nil, fmt.Errorf("invalid ETH amount %q", value)
		}
		wholeWei.Add(wholeWei, fractionWei)
	}
	if wholeWei.Sign() <= 0 {
		return nil, errors.New("ETH amount must be positive")
	}
	return wholeWei, nil
}

func digitsOnly(value string) bool {
	if value == "" {
		return false
	}
	for _, r := range value {
		if !unicode.IsDigit(r) || r > unicode.MaxASCII {
			return false
		}
	}
	return true
}
