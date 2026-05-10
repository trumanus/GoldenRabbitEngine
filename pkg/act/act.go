// Package act campiona azioni a ∈ 𝒜 con costo C(a) ⪯ K (capitolo 14, Volume 2).
package act

import (
	"github.com/trumanus/GoldenRabbitEngine/pkg/budget"
)

// Action è un’azione candidata con costo vettoriale stimato.
type Action struct {
	Name string
	Cost []float64 // stesso ordine del Ledger K
}

// Select sceglie azione ammissibile con perdita L proxy minima (greedy laboratorio).
func Select(ledger *budget.Ledger, candidates []Action, loss []float64) (Action, bool) {
	if ledger == nil || len(candidates) != len(loss) {
		return Action{}, false
	}
	bestIdx := -1
	var best Action
	bestL := 0.0
	first := true
	for i, a := range candidates {
		if len(a.Cost) != len(ledger.K) {
			continue
		}
		if !budget.Feasible(ledger.K, a.Cost) {
			continue
		}
		if first || loss[i] < bestL {
			first = false
			bestL = loss[i]
			bestIdx = i
			best = a
		}
	}
	if bestIdx < 0 {
		return Action{}, false
	}
	return best, true
}
