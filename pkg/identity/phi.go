// Package identity osserva Φ(X) surrogato per drift control (capitoli 7–14, Volume 2).
package identity

import (
	"fmt"

	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
)

// Observable è una vista compressa di identità per telemetria.
type Observable struct {
	TraceCount int
	Eta        float64
	K          float64
	Tick       int64
}

// FromState costruisce osservabile minimale.
func FromState(c state.Cognitive) Observable {
	return Observable{
		TraceCount: len(c.Mem.Traces),
		Eta:        c.Eta,
		K:          c.K,
		Tick:       c.Tick,
	}
}

// DriftScore norma L1 semplice tra due osservabili (ΔΦ laboratorio).
func DriftScore(a, b Observable) float64 {
	d := float64(abs(a.TraceCount - b.TraceCount))
	if a.Eta > b.Eta {
		d += a.Eta - b.Eta
	} else {
		d += b.Eta - a.Eta
	}
	return d
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// String riepilogo runbook notturno.
func (o Observable) String() string {
	return fmt.Sprintf("Phi~ traces=%d eta=%.4f K=%.4f tick=%d", o.TraceCount, o.Eta, o.K, o.Tick)
}
