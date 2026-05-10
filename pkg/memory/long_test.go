package memory

import (
	"testing"

	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
)

func TestDecay(t *testing.T) {
	m := LongMemory{Traces: []Trace{{ID: "a", Weight: 10}}}
	if err := m.ApplyDecay(0.5, []float64{1}); err != nil {
		t.Fatal(err)
	}
	if m.Traces[0].Weight != 6 {
		t.Fatalf("expected 6, got %v", m.Traces[0].Weight)
	}
}

func TestMassCap(t *testing.T) {
	reg := registry.DefaultDocument()
	reg.MaxTraceMass = 5
	m := LongMemory{}
	err := m.Append(reg, Trace{ID: "x", Weight: 10})
	if err == nil {
		t.Fatal("expected mass error")
	}
}
