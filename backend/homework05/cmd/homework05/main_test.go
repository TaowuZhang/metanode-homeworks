package main

import (
	"strings"
	"testing"
)

func TestValidHash(t *testing.T) {
	t.Parallel()

	valid := "0x" + strings.Repeat("a", 64)
	if !validHash(valid) {
		t.Fatalf("validHash(%q) = false, want true", valid)
	}
	if validHash("0x1234") {
		t.Fatal("short hash unexpectedly accepted")
	}
	if validHash("0x" + "z" + strings.Repeat("0", 63)) {
		t.Fatal("non-hex hash unexpectedly accepted")
	}
}

func TestParseAddress(t *testing.T) {
	t.Parallel()

	const valid = "0x000000000000000000000000000000000000dEaD"
	address, err := parseAddress(valid)
	if err != nil {
		t.Fatalf("parseAddress(%q): %v", valid, err)
	}
	if address.Hex() != valid {
		t.Fatalf("parseAddress(%q) = %s", valid, address.Hex())
	}
	if _, err := parseAddress("not-an-address"); err == nil {
		t.Fatal("invalid address unexpectedly accepted")
	}
}
