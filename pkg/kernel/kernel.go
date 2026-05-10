// Package kernel orchestra moduli Golden Rabbit in un tick computazionale (Volume 2 — arcipelago).
package kernel

import (
	"fmt"

	"github.com/trumanus/GoldenRabbitEngine/pkg/budget"
	"github.com/trumanus/GoldenRabbitEngine/pkg/forensic"
	"github.com/trumanus/GoldenRabbitEngine/pkg/homeostasis"
	"github.com/trumanus/GoldenRabbitEngine/pkg/memory"
	"github.com/trumanus/GoldenRabbitEngine/pkg/modulator"
	"github.com/trumanus/GoldenRabbitEngine/pkg/outbox"
	"github.com/trumanus/GoldenRabbitEngine/pkg/promotion"
	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
	"github.com/trumanus/GoldenRabbitEngine/pkg/values"
)

// System legame tra stato cognitivo, economia interna, omeostasi e persistenza ausiliaria.
type System struct {
	Cognitive *state.Cognitive
	Ledger    *budget.Ledger
	Homeo     *homeostasis.State
	Lambda    *modulator.State
	Tau       values.Tensions

	Promotion *promotion.Engine
	Forensic  *forensic.Log
	Outbox    *outbox.Store
	Multi     *memory.MultiStore
}

// NewSystem costruisce sistema coerente con tre canali budget standard.
func NewSystem(reg registry.Document, regHash string, cogn state.Cognitive, ledger *budget.Ledger) *System {
	c := cogn
	return &System{
		Cognitive: &c,
		Ledger:    ledger,
		Homeo:     homeostasis.NewPID(0, 0.05, 0.01, 0.02),
		Tau:       values.Tensions{},
		Promotion: promotion.NewEngine(),
		Forensic:  forensic.NewLog(),
		Outbox:    outbox.New(nil),
		Multi: &memory.MultiStore{
			Scene: memory.ShortBuffer{Max: 64},
			Long:  cogn.Mem,
		},
	}
}

// SyncPromotionBudget copia K_promotion nel campo scalare usato da promotion.Engine.
func (s *System) SyncPromotionBudget() {
	if s.Ledger == nil || s.Cognitive == nil {
		return
	}
	s.Cognitive.K = s.Ledger.PromotionScalar()
}

// InitLambda opzionale (nil-safe).
func (s *System) InitLambda(init, target []float64, alpha float64) error {
	st, err := modulator.New(init, target, alpha)
	if err != nil {
		return err
	}
	s.Lambda = st
	return nil
}

// Tick esegue spend κ, PID su E osservato, λ lenti, ricarica K — poi sincronizza promotion budget.
func (s *System) Tick(observedE float64, kappa []float64) error {
	if s.Ledger == nil || s.Homeo == nil || s.Cognitive == nil {
		return fmt.Errorf("kernel: incomplete system")
	}
	if len(kappa) != len(s.Ledger.K) {
		return fmt.Errorf("kernel: kappa dim %d vs K dim %d", len(kappa), len(s.Ledger.K))
	}
	if err := s.Ledger.Spend(kappa); err != nil {
		return err
	}
	u := s.Homeo.Step(observedE)
	for i := range s.Ledger.Rho {
		s.Ledger.Rho[i] *= (1 + 0.05*u)
		if s.Ledger.Rho[i] < 0 {
			s.Ledger.Rho[i] = 0
		}
	}
	if s.Lambda != nil {
		s.Lambda.Step()
	}
	if err := s.Ledger.EndTick(); err != nil {
		return err
	}
	s.SyncPromotionBudget()
	s.Cognitive.Tick++
	s.Cognitive.Mem = s.Multi.Long
	return nil
}
