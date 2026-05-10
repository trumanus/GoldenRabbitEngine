package kernel

import (
	"testing"

	"github.com/trumanus/GoldenRabbitEngine/pkg/budget"
	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
)

func TestTickAdvances(t *testing.T) {
	reg := registry.DefaultDocument()
	h, _ := registry.HashRegistry(reg)
	c := state.NewCognitive(reg, h, 10, registry.ThetaRaw{})
	led, err := budget.NewLedger([]float64{10, 10, 10})
	if err != nil {
		t.Fatal(err)
	}
	sys := NewSystem(reg, h, c, led)
	sys.SyncPromotionBudget()
	if err := sys.Tick(0.5, []float64{0.5, 0.5, 1}); err != nil {
		t.Fatal(err)
	}
	if sys.Cognitive.Tick != 1 {
		t.Fatalf("tick %d", sys.Cognitive.Tick)
	}
}
