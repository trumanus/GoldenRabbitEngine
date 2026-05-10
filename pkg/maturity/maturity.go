// Package maturity implementa punteggio M di maturità implementativa (capitolo 15, Volume 2).
package maturity

// Flags bit osservabili per audit interno.
type Flags struct {
	KMeasured      bool
	PiSigned       bool
	ReplayCI       bool
	TenantIsolated bool
}

// Weights α_i — devono riflettere dominio; default equiprobabile laboratorio.
type Weights struct {
	A1, A2, A3, A4 float64
}

// DefaultWeights equal 1.0 each (somma 4 — normalizzare in dashboard esterna).
func DefaultWeights() Weights {
	return Weights{A1: 1, A2: 1, A3: 1, A4: 1}
}

// Score calcola M = Σ α_i 𝟙{flag_i}.
func Score(w Weights, f Flags) float64 {
	var m float64
	if f.KMeasured {
		m += w.A1
	}
	if f.PiSigned {
		m += w.A2
	}
	if f.ReplayCI {
		m += w.A3
	}
	if f.TenantIsolated {
		m += w.A4
	}
	return m
}
