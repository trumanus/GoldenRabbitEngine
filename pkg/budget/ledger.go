// Package budget implementa K(t) vettoriale e ammissibilità Σκ ⪯ K (capitolo 3, Volume 2).
package budget

import (
	"errors"
	"fmt"
)

var (
	// ErrInsufficient se una richiesta supera un canale di budget.
	ErrInsufficient = errors.New("budget: insufficient funds on channel")
	// ErrDimensionMismatch se slice hanno lunghezze incoerenti.
	ErrDimensionMismatch = errors.New("budget: dimension mismatch")
)

const (
	ChannelRetrieval = 0
	ChannelACT       = 1
	ChannelPromotion = 2
)

// DefaultChannelNames ordine canonico laboratorio: retrieval, ACT, promozione.
var DefaultChannelNames = []string{"retrieval", "act", "promotion"}

// Ledger tiene K_ℓ, ricariche ρ_ℓ, perdita γ_K e clip (capitolo 3).
type Ledger struct {
	K      []float64
	GammaK []float64
	Rho    []float64
	KMin   []float64
	KMax   []float64
}

// NewLedger tre canali con clip simmetrici.
func NewLedger(k0 []float64) (*Ledger, error) {
	n := len(k0)
	if n == 0 {
		return nil, fmt.Errorf("budget: empty K")
	}
	g := make([]float64, n)
	r := make([]float64, n)
	lo := make([]float64, n)
	hi := make([]float64, n)
	k := make([]float64, n)
	copy(k, k0)
	for i := range k {
		g[i] = 0.05
		r[i] = 2
		lo[i] = 0
		hi[i] = 1e9
	}
	return &Ledger{K: k, GammaK: g, Rho: r, KMin: lo, KMax: hi}, nil
}

// Feasible verifica Σ κ_j ⪯ K componente per componente.
func Feasible(K, kappa []float64) bool {
	if len(K) != len(kappa) {
		return false
	}
	for i := range K {
		if kappa[i] > K[i]+1e-12 {
			return false
		}
	}
	return true
}

// Spend scala immediatamente K del costo (dentro il tick, prima dell’aggiornamento omeostatico).
func (l *Ledger) Spend(kappa []float64) error {
	if len(kappa) != len(l.K) {
		return ErrDimensionMismatch
	}
	if !Feasible(l.K, kappa) {
		return ErrInsufficient
	}
	for i := range l.K {
		l.K[i] -= kappa[i]
		if l.K[i] < 0 {
			l.K[i] = 0
		}
	}
	return nil
}

// EndTick applica K(t+1) = clip((1-γ_K)K + ρ - κ_already_applied_via_Spend è già stato sottratto).
// Qui modelliamo la ricorsione del libro come: dopo le operazioni, ricarica smooth su ciò che resta:
//
//	K_next = clip((1-γ)*K_current + ρ)
//
// (κ è già stato detratti da Spend nel tick corrente).
func (l *Ledger) EndTick() error {
	if len(l.K) != len(l.GammaK) || len(l.K) != len(l.Rho) {
		return ErrDimensionMismatch
	}
	for i := range l.K {
		next := (1-l.GammaK[i])*l.K[i] + l.Rho[i]
		l.K[i] = clamp(next, l.KMin[i], l.KMax[i])
	}
	return nil
}

// PromotionScalar esposto per integrazione con promotion.Engine che usa uno scalare K.
func (l *Ledger) PromotionScalar() float64 {
	if len(l.K) <= ChannelPromotion {
		return 0
	}
	return l.K[ChannelPromotion]
}

// SetPromotionScalar sincronizza il canale promozione (helper kernel).
func (l *Ledger) SetPromotionScalar(v float64) {
	if len(l.K) > ChannelPromotion {
		l.K[ChannelPromotion] = v
	}
}

func clamp(x, lo, hi float64) float64 {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}
