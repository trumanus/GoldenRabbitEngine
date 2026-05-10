// Package homeostasis modella E, ε_E, d e PID su comando u (capitolo 4, Volume 2).
package homeostasis

import (
	"math"
)

// State tiene tensione progettuale E rispetto a E★ e una densità costruttiva scalare (primo modo di d⃗).
type State struct {
	E       float64
	EStar   float64
	Density float64 // componente sintetica di d⃗ per dashboard

	KP float64
	KI float64
	KD float64

	integral   float64
	prevEpsilon float64
	u0         float64
}

// NewPID crea regolatore con coefficienti clip‑friendly.
func NewPID(eStar, kp, ki, kd float64) *State {
	return &State{
		EStar: eStar,
		KP:    kp,
		KI:    ki,
		KD:    kd,
		u0:    1,
	}
}

// Epsilon è ε_E(t)=E(t)-E★.
func (s *State) Epsilon() float64 {
	return s.E - s.EStar
}

// Step aggiorna stato dopo osservazione di E(t) e integra PID discreto (Figura 4‑1).
func (s *State) Step(observedE float64) float64 {
	s.E = observedE
	eps := s.Epsilon()
	s.integral += eps
	de := eps - s.prevEpsilon
	u := s.u0 - s.KP*eps - s.KI*s.integral - s.KD*de
	s.prevEpsilon = eps
	return u
}

// DensityInBand verifica d_min ≤ d ≤ d_max (vettore ridotto a scalare in questo prototipo).
func DensityInBand(d, dMin, dMax float64) bool {
	return d >= dMin && d <= dMax
}

// RecoverHint stima se |ε_E| ≤ δ per H passi consecutivi (callback per serie storica esterna).
func RecoverHint(epsilon, delta float64) bool {
	return math.Abs(epsilon) <= delta
}
