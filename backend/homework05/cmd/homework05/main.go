package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/TaowuZhang/metanode-homeworks/backend/homework05/counter"
	"github.com/TaowuZhang/metanode-homeworks/backend/homework05/ethrpc"
	"github.com/ethereum/go-ethereum/common"
)

const commandTimeout = 3 * time.Minute

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		printUsage()
		return errors.New("command is required")
	}

	command := os.Args[1]
	switch command {
	case "help", "-h", "--help":
		printUsage()
		return nil
	case "block", "tx", "transfer", "counter-deploy", "counter-get", "counter-inc":
		// Commands below require an HTTP/HTTPS RPC connection.
	default:
		printUsage()
		return fmt.Errorf("unknown command %q", command)
	}

	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	rpcURL := strings.TrimSpace(os.Getenv("SEPOLIA_RPC_URL"))
	client, err := ethrpc.Dial(ctx, rpcURL)
	if err != nil {
		return err
	}
	defer client.Close()

	switch command {
	case "block":
		return runBlock(ctx, client, os.Args[2:])
	case "tx":
		return runTransaction(ctx, client, os.Args[2:])
	case "transfer":
		return runTransfer(ctx, client, os.Args[2:])
	case "counter-deploy":
		return runCounterDeploy(ctx, client, os.Args[2:])
	case "counter-get":
		return runCounterGet(ctx, client, os.Args[2:])
	case "counter-inc":
		return runCounterIncrement(ctx, client, os.Args[2:])
	default:
		return fmt.Errorf("unhandled command %q", command)
	}
}

func runBlock(ctx context.Context, client *ethrpc.Client, args []string) error {
	if len(args) > 1 {
		return errors.New("usage: block [number|latest]")
	}
	var number *big.Int
	if len(args) == 1 && args[0] != "latest" {
		parsed, ok := new(big.Int).SetString(args[0], 10)
		if !ok || parsed.Sign() < 0 {
			return fmt.Errorf("invalid block number %q", args[0])
		}
		number = parsed
	}

	info, err := client.Block(ctx, number)
	if err != nil {
		return err
	}
	fmt.Printf("number: %s\n", info.Number)
	fmt.Printf("hash: %s\n", info.Hash.Hex())
	fmt.Printf("parent: %s\n", info.ParentHash.Hex())
	fmt.Printf("timestamp: %s\n", time.Unix(int64(info.Timestamp), 0).UTC().Format(time.RFC3339))
	fmt.Printf("transactions: %d\n", info.Transactions)
	return nil
}

func runTransaction(ctx context.Context, client *ethrpc.Client, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: tx <transaction-hash>")
	}
	if !validHash(args[0]) {
		return fmt.Errorf("invalid transaction hash %q", args[0])
	}
	info, err := client.Transaction(ctx, common.HexToHash(args[0]))
	if err != nil {
		return err
	}
	fmt.Printf("hash: %s\n", info.Hash.Hex())
	fmt.Printf("from: %s\n", info.From.Hex())
	if info.To == nil {
		fmt.Println("to: <contract creation>")
	} else {
		fmt.Printf("to: %s\n", info.To.Hex())
	}
	fmt.Printf("valueWei: %s\n", info.Value)
	fmt.Printf("nonce: %d\n", info.Nonce)
	fmt.Printf("gas: %d\n", info.Gas)
	fmt.Printf("type: %d\n", info.Type)
	fmt.Printf("gasPriceWei: %s\n", info.GasPrice)
	fmt.Printf("gasTipCapWei: %s\n", info.GasTipCap)
	fmt.Printf("gasFeeCapWei: %s\n", info.GasFeeCap)
	fmt.Printf("pending: %t\n", info.Pending)
	return nil
}

func runTransfer(ctx context.Context, client *ethrpc.Client, args []string) error {
	if len(args) != 2 {
		return errors.New("usage: transfer <to-address> <amount-eth>")
	}
	to, err := parseAddress(args[0])
	if err != nil {
		return err
	}
	amount, err := ethrpc.ParseETH(args[1])
	if err != nil {
		return err
	}
	key := os.Getenv("SEPOLIA_PRIVATE_KEY")
	hash, from, err := client.SendETH(ctx, key, to, amount)
	if err != nil {
		return err
	}
	fmt.Printf("from: %s\n", from.Hex())
	fmt.Printf("to: %s\n", to.Hex())
	fmt.Printf("amountWei: %s\n", amount)
	fmt.Printf("txHash: %s\n", hash.Hex())
	return nil
}

func runCounterDeploy(ctx context.Context, client *ethrpc.Client, args []string) error {
	if len(args) != 0 {
		return errors.New("usage: counter-deploy")
	}
	address, result, err := counter.Deploy(ctx, client, os.Getenv("SEPOLIA_PRIVATE_KEY"))
	if err != nil {
		return err
	}
	fmt.Printf("contract: %s\n", address.Hex())
	fmt.Printf("txHash: %s\n", result.Hash.Hex())
	fmt.Printf("block: %s\n", result.BlockNumber)
	return nil
}

func runCounterGet(ctx context.Context, client *ethrpc.Client, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: counter-get <contract-address>")
	}
	address, err := parseAddress(args[0])
	if err != nil {
		return err
	}
	value, err := counter.Number(ctx, client, address)
	if err != nil {
		return err
	}
	fmt.Printf("number: %s\n", value)
	return nil
}

func runCounterIncrement(ctx context.Context, client *ethrpc.Client, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: counter-inc <contract-address>")
	}
	address, err := parseAddress(args[0])
	if err != nil {
		return err
	}
	result, err := counter.Increment(ctx, client, os.Getenv("SEPOLIA_PRIVATE_KEY"), address)
	if err != nil {
		return err
	}
	fmt.Printf("txHash: %s\n", result.Hash.Hex())
	fmt.Printf("block: %s\n", result.BlockNumber)
	return nil
}

func parseAddress(value string) (common.Address, error) {
	if !common.IsHexAddress(value) {
		return common.Address{}, fmt.Errorf("invalid Ethereum address %q", value)
	}
	return common.HexToAddress(value), nil
}

func validHash(value string) bool {
	value = strings.TrimPrefix(value, "0x")
	if len(value) != 64 {
		return false
	}
	for _, r := range value {
		if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
			return false
		}
	}
	return true
}

func printUsage() {
	fmt.Print(`MetaNode backend homework05 - Sepolia with go-ethereum

Environment:
  SEPOLIA_RPC_URL       Sepolia HTTP/HTTPS JSON-RPC endpoint
  SEPOLIA_PRIVATE_KEY   test-account private key (write commands only)

Commands:
  block [number|latest]             query a block
  tx <transaction-hash>             query a transaction
  transfer <to-address> <amount>    send ETH, amount is in ETH
  counter-deploy                    deploy Counter.sol
  counter-get <contract-address>    read Counter.number
  counter-inc <contract-address>    call Counter.increment
`)
}
