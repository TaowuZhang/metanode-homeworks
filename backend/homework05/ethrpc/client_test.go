package ethrpc

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestParseETH(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "one eth", input: "1", want: "1000000000000000000"},
		{name: "milliether", input: "0.001", want: "1000000000000000"},
		{name: "one wei", input: "0.000000000000000001", want: "1"},
		{name: "leading dot", input: ".5", want: "500000000000000000"},
		{name: "too precise", input: "0.0000000000000000001", wantErr: true},
		{name: "zero", input: "0", wantErr: true},
		{name: "negative", input: "-1", wantErr: true},
		{name: "invalid", input: "abc", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ParseETH(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseETH(%q) expected error", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseETH(%q): %v", tt.input, err)
			}
			want, ok := new(big.Int).SetString(tt.want, 10)
			if !ok {
				t.Fatalf("invalid test expectation %q", tt.want)
			}
			if got.Cmp(want) != 0 {
				t.Fatalf("ParseETH(%q) = %s, want %s", tt.input, got, want)
			}
		})
	}
}

func TestPrivateKeyFromHexAcceptsPrefix(t *testing.T) {
	t.Parallel()

	generated, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generate test key: %v", err)
	}
	key := hex.EncodeToString(crypto.FromECDSA(generated))

	withoutPrefix, err := PrivateKeyFromHex(key)
	if err != nil {
		t.Fatalf("parse key without prefix: %v", err)
	}
	withPrefix, err := PrivateKeyFromHex("0x" + key)
	if err != nil {
		t.Fatalf("parse key with prefix: %v", err)
	}
	if withoutPrefix.D.Cmp(withPrefix.D) != 0 {
		t.Fatal("0x prefix changed parsed private key")
	}
}

func TestRequireSepolia(t *testing.T) {
	t.Parallel()

	if err := requireSepolia(big.NewInt(SepoliaChainID)); err != nil {
		t.Fatalf("Sepolia rejected: %v", err)
	}
	if err := requireSepolia(big.NewInt(1)); err == nil {
		t.Fatal("mainnet chain ID unexpectedly accepted")
	}
	if err := requireSepolia(nil); err == nil {
		t.Fatal("nil chain ID unexpectedly accepted")
	}
}

func TestNewDynamicFeeTransfer(t *testing.T) {
	t.Parallel()

	chainID := big.NewInt(SepoliaChainID)
	to := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	value := big.NewInt(1_000_000_000_000_000)
	tip := big.NewInt(2_000_000_000)
	fee := big.NewInt(42_000_000_000)
	tx := newDynamicFeeTransfer(chainID, 7, to, value, tip, fee)

	if tx.Type() != types.DynamicFeeTxType {
		t.Fatalf("tx type = %d, want %d", tx.Type(), types.DynamicFeeTxType)
	}
	if tx.ChainId().Cmp(chainID) != 0 {
		t.Fatalf("chain ID = %s, want %s", tx.ChainId(), chainID)
	}
	if tx.Nonce() != 7 || tx.Gas() != 21_000 {
		t.Fatalf("nonce/gas = %d/%d, want 7/21000", tx.Nonce(), tx.Gas())
	}
	if tx.To() == nil || *tx.To() != to {
		t.Fatalf("to = %v, want %s", tx.To(), to.Hex())
	}
	if tx.Value().Cmp(value) != 0 || tx.GasTipCap().Cmp(tip) != 0 || tx.GasFeeCap().Cmp(fee) != 0 {
		t.Fatal("dynamic fee transaction fields do not match inputs")
	}
}
