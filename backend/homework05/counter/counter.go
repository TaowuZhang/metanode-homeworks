package counter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/TaowuZhang/metanode-homeworks/backend/homework05/bindings"
	"github.com/TaowuZhang/metanode-homeworks/backend/homework05/ethrpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type TxResult struct {
	Hash        common.Hash
	BlockNumber *big.Int
}

func Deploy(ctx context.Context, client *ethrpc.Client, privateKeyHex string) (common.Address, *TxResult, error) {
	auth, _, err := client.NewTransactor(ctx, privateKeyHex)
	if err != nil {
		return common.Address{}, nil, err
	}

	address, tx, _, err := bindings.DeployCounter(auth, client.RPC())
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("deploy Counter: %w", err)
	}
	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, &TxResult{Hash: tx.Hash(), BlockNumber: receipt.BlockNumber}, nil
}

func Number(ctx context.Context, client *ethrpc.Client, address common.Address) (*big.Int, error) {
	contract, err := bindings.NewCounter(address, client.RPC())
	if err != nil {
		return nil, fmt.Errorf("bind Counter: %w", err)
	}
	value, err := contract.Number(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("read Counter.number: %w", err)
	}
	return value, nil
}

func Increment(ctx context.Context, client *ethrpc.Client, privateKeyHex string, address common.Address) (*TxResult, error) {
	contract, err := bindings.NewCounter(address, client.RPC())
	if err != nil {
		return nil, fmt.Errorf("bind Counter: %w", err)
	}
	auth, _, err := client.NewTransactor(ctx, privateKeyHex)
	if err != nil {
		return nil, err
	}
	tx, err := contract.Increment(auth)
	if err != nil {
		return nil, fmt.Errorf("send increment: %w", err)
	}
	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return nil, err
	}
	return &TxResult{Hash: tx.Hash(), BlockNumber: receipt.BlockNumber}, nil
}
