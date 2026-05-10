// Package conflict risolve ambiguità tra candidati ACT/policy usando τ (capitolo 10, Volume 2).
package conflict

import (
	"github.com/trumanus/GoldenRabbitEngine/pkg/values"
)

// Candidate è un’azione con tag di dominio per pesatura τ.
type Candidate struct {
	Name   string
	Tags   []string
	Loss   float64 // L locale (da minimizzare)
	Risk   float64 // R stimato (κ rischio / penality)
}

// Resolution esito della risoluzione.
type Resolution struct {
	Chosen Candidate
	Reason string
	Score  float64
}

// Resolve sceglie il candidato che minimizza Loss + Σ τ_tag * Risk (schema leggero).
func Resolve(tau values.Tensions, candidates []Candidate) (Resolution, bool) {
	if len(candidates) == 0 {
		return Resolution{}, false
	}
	best := candidates[0]
	bestScore := score(tau, best)
	for _, c := range candidates[1:] {
		s := score(tau, c)
		if s < bestScore {
			bestScore = s
			best = c
		}
	}
	return Resolution{
		Chosen: best,
		Reason: "min weighted loss+risk under tau",
		Score:  bestScore,
	}, true
}

func score(tau values.Tensions, c Candidate) float64 {
	w := 1.0
	for _, tag := range c.Tags {
		w += tau.Weight(tag)
	}
	return c.Loss + w*c.Risk
}
