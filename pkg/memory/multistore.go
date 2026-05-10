package memory

import (
	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
)

// MultiStore combina scena corta e lunga cognitiva con masse per layer (capitolo 8).
type MultiStore struct {
	Scene ShortBuffer
	Long  LongMemory
}

// MassByLayer somma pesi delle tracce etichettate per layer.
func (m MultiStore) MassByLayer() map[Layer]float64 {
	out := map[Layer]float64{}
	for _, t := range m.Long.Traces {
		ell := Layer(t.Layer)
		if ell == "" {
			ell = LayerCognitive
		}
		out[ell] += t.Weight
	}
	return out
}

// AppendCognitive promuove nella lunga con layer cognitivo esplicito.
func (m *MultiStore) AppendCognitive(reg registry.Document, t Trace) error {
	t.Layer = string(LayerCognitive)
	return m.Long.Append(reg, t)
}
