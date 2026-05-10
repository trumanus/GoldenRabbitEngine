package memory

import (
	"math"
)

// RetrievalWeights coefficienti (α,β,γ,δ) del capitolo 8 — vivono nel Registry in produzione.
type RetrievalWeights struct {
	AlphaSim float64
	BetaRisk float64
	GammaAge float64
	DeltaViol float64
}

// ScoreItem contributi vettoriali allo score s_i (auditabile).
type ScoreItem struct {
	SimTerm      float64
	RiskTerm     float64
	AgeTerm      float64
	ViolationTerm float64
	Total        float64
}

// CompositeScore calcola s_i = α·sim − β·R − γ·age − δ·Viol (schema libro).
func CompositeScore(sim, risk, age, viol float64, w RetrievalWeights) ScoreItem {
	s := ScoreItem{
		SimTerm:       w.AlphaSim * sim,
		RiskTerm:      -w.BetaRisk * risk,
		AgeTerm:       -w.GammaAge * age,
		ViolationTerm: -w.DeltaViol * viol,
	}
	s.Total = s.SimTerm + s.RiskTerm + s.AgeTerm + s.ViolationTerm
	return s
}

// NormalizeCosine stub similarity ∈ [0,1] da prodotto scalare già normalizzato.
func NormalizeCosine(dot float64) float64 {
	return math.Max(0, math.Min(1, dot))
}
