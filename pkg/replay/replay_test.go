package replay

import (
	"encoding/json"
	"testing"

	"github.com/trumanus/GoldenRabbitEngine/pkg/event"
	"github.com/trumanus/GoldenRabbitEngine/pkg/promotion"
	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
)

func TestReplayParitySnapshots(t *testing.T) {
	reg := registry.DefaultDocument()
	h, err := registry.HashRegistry(reg)
	if err != nil {
		t.Fatal(err)
	}
	raw := registry.ThetaRaw{SalienceGate: 0.4, Eta: 0.05}

	runOnce := func() (string, error) {
		c := state.NewCognitive(reg, h, 100, raw)
		eng := promotion.NewEngine()
		ev := event.Typed{Type: "order.confirmed", Attrs: map[string]string{"tenant": "t", "order_id": "o1"}}
		steps := []Step{
			{
				Tick: 1,
				Intent: promotion.Intent{
					IdempotencyKey:       "idem-1",
					ThetaVersion:         "0.0.1",
					RegistryHashExpected: h,
					KappaPi:              3,
				},
				Delta: func(c *state.Cognitive) error {
					return promotion.AppendTraceFromEvent(c, ev, "tr-1", 2, "order")
				},
			},
		}
		if err := Run(&c, eng, steps); err != nil {
			return "", err
		}
		snap, err := c.Snapshot()
		if err != nil {
			return "", err
		}
		b, err := json.Marshal(snap)
		return string(b), err
	}

	a, err := runOnce()
	if err != nil {
		t.Fatal(err)
	}
	b, err := runOnce()
	if err != nil {
		t.Fatal(err)
	}
	if a != b {
		t.Fatalf("replay parity failed:\n%s\nvs\n%s", a, b)
	}
}
