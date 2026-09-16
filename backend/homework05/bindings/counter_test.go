package bindings

import "testing"

func TestCounterMetadata(t *testing.T) {
	t.Parallel()

	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		t.Fatalf("parse ABI: %v", err)
	}
	if parsed == nil {
		t.Fatal("parsed ABI is nil")
	}
	for _, name := range []string{"number", "increment"} {
		if _, ok := parsed.Methods[name]; !ok {
			t.Fatalf("ABI missing method %q", name)
		}
	}
	if _, ok := parsed.Events["Incremented"]; !ok {
		t.Fatal("ABI missing Incremented event")
	}
	if CounterBin == "" || CounterBin == "0x" {
		t.Fatal("Counter bytecode is empty")
	}
}
