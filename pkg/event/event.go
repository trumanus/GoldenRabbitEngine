// Package event implementa la catena x → z → e ∈ ℰ ∪ {⊥} (capitolo 2, Volume 2).
package event

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
)

var (
	// ErrBot indica che il materiale non è ancora un evento tipizzato (⊥).
	ErrBot = errors.New("event: not yet typed (bot)")
)

// RawInput è x(t) al confine — payload grezzo osservabile / attestabile.
type RawInput struct {
	RequestID    string          `json:"request_id"`
	ReceivedAt   time.Time       `json:"received_at"`
	Tenant       string          `json:"tenant"`
	Payload      json.RawMessage `json:"payload"`
	Correlation  string          `json:"correlation_id,omitempty"`
	Idempotency  string          `json:"idempotency_key,omitempty"`
	DeclaredTick int64           `json:"project_tick"` // tick di progetto n
}

// Encoded è z(t) = encode(x) — vista interna pre-decisionale.
type Encoded struct {
	RequestID    string             `json:"request_id"`
	Tenant       string             `json:"tenant"`
	Features     map[string]float64 `json:"features,omitempty"` // stub numeriche / embedding surrogate
	RedactedHint string             `json:"redacted_hint,omitempty"`
	PayloadRef   string             `json:"payload_ref,omitempty"` // hash o URI verso storage immutabile
	Correlation  string             `json:"correlation_id,omitempty"`
	Idempotency  string             `json:"idempotency_key,omitempty"`
	ProjectTick  int64              `json:"project_tick"`
}

// Typed è e(t) ∈ ℰ dopo contratto Registry.
type Typed struct {
	Type           string            `json:"type"`
	Attrs          map[string]string `json:"attrs"`
	PayloadRef     string            `json:"payload_ref,omitempty"`
	CorrelationID  string            `json:"correlation_id"`
	IdempotencyKey string            `json:"idempotency_key"`
	ProjectTick    int64             `json:"project_tick"`
	OccurredAt     time.Time         `json:"occurred_at"`
}

// Encoder traduce ingresso grezzo in z (deterministico a parità di policy).
type Encoder interface {
	Encode(in RawInput) (Encoded, error)
}

// FunctionalEncoder adatta una funzione pura all'interfaccia Encoder.
type FunctionalEncoder func(RawInput) (Encoded, error)

// Encode implementa Encoder.
func (f FunctionalEncoder) Encode(in RawInput) (Encoded, error) { return f(in) }

// Pipeline incapsula type + validate verso ℰ.
type Pipeline struct {
	Reg registry.Document
	Enc Encoder
}

// FromRaw esegue encode + tipizzazione; fallisce con ErrBot se σ < σ* o tipo assente.
func (p Pipeline) FromRaw(in RawInput, typ string, attrs map[string]string, sigma float64) (Typed, error) {
	if sigma < p.Reg.SigmaStar {
		return Typed{}, ErrBot
	}
	z, err := p.Enc.Encode(in)
	if err != nil {
		return Typed{}, err
	}
	if typ == "" {
		return Typed{}, ErrBot
	}
	if err := p.Reg.ValidateAttrs(typ, attrs); err != nil {
		return Typed{}, err
	}
	idem := in.Idempotency
	if idem == "" {
		idem = in.RequestID
	}
	corr := in.Correlation
	if corr == "" {
		corr = in.RequestID
	}
	return Typed{
		Type:           typ,
		Attrs:          attrs,
		PayloadRef:     z.PayloadRef,
		CorrelationID:  corr,
		IdempotencyKey: idem,
		ProjectTick:    z.ProjectTick,
		OccurredAt:     time.Now().UTC(),
	}, nil
}

// MetricsRatio ρ_evt semplicistico per laboratorio (capitolo 2).
type MetricsRatio struct {
	TypedAttempts   int64
	BotOrRejected   int64
	SchemaRejected  int64
	RawIngressCount int64
}

// RhoEvt rapporto tipizzati su ingressi grezzi.
func (m MetricsRatio) RhoEvt() float64 {
	if m.RawIngressCount == 0 {
		return 0
	}
	return float64(m.TypedAttempts) / float64(m.RawIngressCount)
}
