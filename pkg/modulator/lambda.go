// Package modulator modella λ lenti (vento endocrino computazionale — capitolo 5 / 14, Volume 2).
package modulator

import (
	"fmt"
)

// State è un vettore λ con aggiornamento lento verso obiettivo.
type State struct {
	Lambda []float64
	Target []float64
	Alpha  float64 // passo (0,1]
}

// New crea modulatori con stessa dimensione.
func New(init, target []float64, alpha float64) (*State, error) {
	if len(init) != len(target) || len(init) == 0 {
		return nil, fmt.Errorf("modulator: dimension mismatch")
	}
	if alpha <= 0 || alpha > 1 {
		return nil, fmt.Errorf("modulator: alpha out of (0,1]")
	}
	l := make([]float64, len(init))
	copy(l, init)
	t := make([]float64, len(target))
	copy(t, target)
	return &State{Lambda: l, Target: t, Alpha: alpha}, nil
}

// Step λ ← (1-α)λ + α target (schema esponenziale semplice).
func (s *State) Step() {
	for i := range s.Lambda {
		s.Lambda[i] = (1-s.Alpha)*s.Lambda[i] + s.Alpha*s.Target[i]
	}
}
