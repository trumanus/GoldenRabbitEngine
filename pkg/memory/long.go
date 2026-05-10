// Package memory implementa tracce in 𝒟𝓂 𝓁 con decadimento (capitolo 2 e 14, Volume 2).
package memory

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
)

var (
	// ErrMassExceeded quando Σ w supera il budget pubblicato nel Registry.
	ErrMassExceeded = errors.New("memory: total trace mass exceeds constitutional bound")
)

// Trace è una traccia i con peso w_i ≥ 0 e tipo nominabile.
type Trace struct {
	ID       string             `json:"id"`
	Class    string             `json:"class"`
	Layer    string             `json:"layer,omitempty"` // ℓ — capitolo 8
	Weight   float64            `json:"weight"`
	Metadata map[string]string  `json:"metadata,omitempty"`
}

// LongMemory aggrega tracce comportamentali promosse (sottoinsieme di 𝒟𝓂 𝓁).
type LongMemory struct {
	Traces []Trace `json:"traces"`
}

// TotalMass somma i pesi — utile per confronto con MaxTraceMass.
func (m LongMemory) TotalMass() float64 {
	var s float64
	for _, t := range m.Traces {
		s += t.Weight
	}
	return s
}

// ApplyDecay applica w_i(t+1) = γ_i w_i(t) + Δ_i (capitolo 2).
func (m *LongMemory) ApplyDecay(gamma float64, delta []float64) error {
	if gamma <= 0 || gamma > 1 {
		return fmt.Errorf("memory: gamma out of (0,1]: %v", gamma)
	}
	if len(delta) > 0 && len(delta) != len(m.Traces) {
		return fmt.Errorf("memory: delta length %d != traces %d", len(delta), len(m.Traces))
	}
	for i := range m.Traces {
		d := 0.0
		if len(delta) > 0 {
			d = delta[i]
		}
		m.Traces[i].Weight = gamma*m.Traces[i].Weight + d
		if m.Traces[i].Weight < 0 {
			m.Traces[i].Weight = 0
		}
	}
	return nil
}

// Append aggiunge una traccia verificando il tetto costituzionale.
func (m *LongMemory) Append(reg registry.Document, t Trace) error {
	next := m.TotalMass() + t.Weight
	if reg.MaxTraceMass > 0 && next > reg.MaxTraceMass {
		return fmt.Errorf("%w: next=%.4f cap=%.4f", ErrMassExceeded, next, reg.MaxTraceMass)
	}
	m.Traces = append(m.Traces, t)
	return nil
}

// Marshal canonico per embedding in state.Snapshot.
func (m LongMemory) MarshalJSONBytes() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal da snapshot.
func UnmarshalJSONBytes(b []byte) (LongMemory, error) {
	var m LongMemory
	err := json.Unmarshal(b, &m)
	return m, err
}
