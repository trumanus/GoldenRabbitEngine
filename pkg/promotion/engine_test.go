package promotion

import (
	"testing"

	"github.com/trumanus/GoldenRabbitEngine/pkg/event"
	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
)

func TestIdempotentApply(t *testing.T) {
	reg := registry.DefaultDocument()
	h, err := registry.HashRegistry(reg)
	if err != nil {
		t.Fatal(err)
	}
	c := state.NewCognitive(reg, h, 10, registry.ThetaRaw{SalienceGate: 0.5, Eta: 0.1})
	eng := NewEngine()
	intent := Intent{
		IdempotencyKey:       "k1",
		ThetaVersion:         "0.0.1",
		RegistryHashExpected: h,
		KappaPi:              2,
	}
	ev := event.Typed{Type: "order.confirmed", Attrs: map[string]string{"tenant": "t"}}
	ap1, err := eng.Apply(&c, intent, func() error {
		return AppendTraceFromEvent(&c, ev, "tr-1", 1, "order")
	})
	if err != nil || !ap1 {
		t.Fatalf("first apply: %v applied=%v", err, ap1)
	}
	ap2, err := eng.Apply(&c, intent, func() error {
		t.Fatal("delta should not run twice")
		return nil
	})
	if err != nil || ap2 {
		t.Fatalf("second apply: %v applied=%v", err, ap2)
	}
	if c.K != 8 {
		t.Fatalf("expected K=8, got %v", c.K)
	}
	if len(c.Mem.Traces) != 1 {
		t.Fatalf("expected 1 trace")
	}
}

func TestBudgetBlocks(t *testing.T) {
	reg := registry.DefaultDocument()
	h, _ := registry.HashRegistry(reg)
	c := state.NewCognitive(reg, h, 1, registry.ThetaRaw{})
	eng := NewEngine()
	_, err := eng.Apply(&c, Intent{IdempotencyKey: "x", KappaPi: 5}, func() error { return nil })
	if err != ErrInsufficientBudget {
		t.Fatalf("expected budget error, got %v", err)
	}
}
