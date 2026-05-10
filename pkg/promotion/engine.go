// Package promotion implementa 𝒥_π con idempotenza su chiavi naturali (capitoli 2, 15, Volume 2).
package promotion

import (
	"errors"
	"fmt"

	"github.com/trumanus/GoldenRabbitEngine/pkg/event"
	"github.com/trumanus/GoldenRabbitEngine/pkg/memory"
	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
)

var (
	// ErrInsufficientBudget quando κ_π supera K disponibile (capitolo 3 — modello semplificato).
	ErrInsufficientBudget = errors.New("promotion: insufficient K budget")
	// ErrRegistryMismatch se hash Registry non coincide con lo stato atteso.
	ErrRegistryMismatch = errors.New("promotion: registry hash mismatch")
)

// Intent identifica l'intento promozionale ripetibile dalla rete (idempotenza).
type Intent struct {
	IdempotencyKey       string
	ThetaVersion         string // versione π_θ / bundle promozionale
	RegistryHashExpected string // hash costituzione al commit
	KappaPi              float64
}

// Engine mantiene l'insieme Applied ⊂ chiavi già eseguite.
type Engine struct {
	Applied map[string]struct{}
}

// NewEngine crea motore con set vuoto.
func NewEngine() *Engine {
	return &Engine{Applied: map[string]struct{}{}}
}

// Apply esegue Π come transizione commitata se k ∉ Applied e budget disponibile:
//
//	X ← X ⊕ Δ(k) se k ∉ Applied, altrimenti X (capitolo 15).
func (e *Engine) Apply(c *state.Cognitive, in Intent, delta func() error) (applied bool, err error) {
	if in.IdempotencyKey == "" {
		return false, fmt.Errorf("promotion: empty idempotency key")
	}
	if _, exists := e.Applied[in.IdempotencyKey]; exists {
		return false, nil
	}
	if in.RegistryHashExpected != "" && in.RegistryHashExpected != c.RegHash {
		return false, fmt.Errorf("%w: want %s got %s", ErrRegistryMismatch, in.RegistryHashExpected, c.RegHash)
	}
	if in.KappaPi > c.K {
		return false, ErrInsufficientBudget
	}
	if err := delta(); err != nil {
		return false, err
	}
	c.K -= in.KappaPi
	e.Applied[in.IdempotencyKey] = struct{}{}
	return true, nil
}

// AppendTraceFromEvent è un helper che materializza Δ minimalista su 𝒟𝓂 𝓁 dopo evento tipizzato.
func AppendTraceFromEvent(c *state.Cognitive, ev event.Typed, traceID string, weight float64, class string) error {
	t := memory.Trace{
		ID:     traceID,
		Class:  class,
		Layer:  string(memory.LayerCognitive),
		Weight: weight,
		Metadata: map[string]string{
			"event_type": ev.Type,
			"tenant":     ev.Attrs["tenant"],
		},
	}
	return c.Mem.Append(c.Reg, t)
}
