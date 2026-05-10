// Package registry modella la costituzione cognitiva (capitolo 7, Volume 2):
// schema degli eventi ℰ, versione semver, hash pubblicabile, bounds su θ per ΠΘ.
package registry

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

var (
	// ErrUnknownEventType indica un tipo non ammesso dal Registry corrente.
	ErrUnknownEventType = errors.New("registry: unknown event type")
	// ErrInvariantViolation segnala payload fuori dai bounds costituzionali.
	ErrInvariantViolation = errors.New("registry: invariant violation")
)

// Document è una fotografia versionata di 𝒞_Reg^(v) — sottoinsieme implementabile
// come artefatto serializzabile e firmabile in pipeline (cap. 15).
type Document struct {
	Version       string            `json:"version"` // semver costituzionale
	EventTypes    map[string]Schema `json:"event_types"`
	ThetaBounds   ThetaBounds       `json:"theta_bounds"`
	SigmaStar     float64           `json:"sigma_star"` // soglia salienza σ* pubblicata
	MaxTraceMass  float64           `json:"max_trace_mass,omitempty"`
	Notes         string            `json:"notes,omitempty"`
}

// Schema descrive vincoli minimi su un tipo e ∈ ℰ.
type Schema struct {
	Description   string   `json:"description,omitempty"`
	RequiredAttrs []string `json:"required_attrs,omitempty"`
}

// ThetaBounds definisce un corridoio Θ^(v) per proiezione ΠΘ (Figura 7‑1, Volume 2).
type ThetaBounds struct {
	SalienceGateMin float64 `json:"salience_gate_min"` // clamp inferiore effetto soglia
	SalienceGateMax float64 `json:"salience_gate_max"`
	EtaMin          float64 `json:"eta_min"`
	EtaMax          float64 `json:"eta_max"`
}

// DefaultDocument costituzione minimale per laboratorio / CI.
func DefaultDocument() Document {
	return Document{
		Version: "0.1.0",
		EventTypes: map[string]Schema{
			"order.confirmed": {RequiredAttrs: []string{"tenant", "order_id"}},
			"preference.promoted": {
				Description:   "Promozione verso lunga comportamentale",
				RequiredAttrs: []string{"tenant", "scope", "rationale_ref"},
			},
		},
		ThetaBounds: ThetaBounds{
			SalienceGateMin: 0,
			SalienceGateMax: 1,
			EtaMin:          -1,
			EtaMax:          1,
		},
		SigmaStar:    0.55,
		MaxTraceMass: 1e6,
	}
}

// HashRegistry produce un digest stabile da allegare a snapshot e runbook (cap. 2, 15).
func HashRegistry(d Document) (string, error) {
	canonical, err := canonicalJSON(d)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(canonical)
	return hex.EncodeToString(sum[:]), nil
}

func canonicalJSON(d Document) ([]byte, error) {
	// Ordine stabile delle chiavi di primo livello + event_types ordinati lessicograficamente.
	type wire struct {
		Version      string            `json:"version"`
		EventTypes   map[string]Schema `json:"event_types"`
		ThetaBounds  ThetaBounds       `json:"theta_bounds"`
		SigmaStar    float64           `json:"sigma_star"`
		MaxTraceMass float64           `json:"max_trace_mass,omitempty"`
		Notes        string            `json:"notes,omitempty"`
	}
	names := make([]string, 0, len(d.EventTypes))
	for k := range d.EventTypes {
		names = append(names, k)
	}
	sort.Strings(names)
	ordered := make(map[string]Schema, len(names))
	for _, k := range names {
		s := d.EventTypes[k]
		sort.Strings(s.RequiredAttrs)
		ordered[k] = s
	}
	w := wire{
		Version:      d.Version,
		EventTypes:   ordered,
		ThetaBounds:  d.ThetaBounds,
		SigmaStar:    d.SigmaStar,
		MaxTraceMass: d.MaxTraceMass,
		Notes:        d.Notes,
	}
	return json.Marshal(w)
}

// ValidateEventType verifica che il tipo sia ammesso da ℰ.
func (d Document) ValidateEventType(typ string) error {
	if typ == "" {
		return fmt.Errorf("%w: empty type", ErrUnknownEventType)
	}
	if _, ok := d.EventTypes[typ]; !ok {
		return fmt.Errorf("%w: %q", ErrUnknownEventType, typ)
	}
	return nil
}

// ValidateAttrs controlla attributi obbligatori dichiarati per il tipo.
func (d Document) ValidateAttrs(typ string, attrs map[string]string) error {
	if err := d.ValidateEventType(typ); err != nil {
		return err
	}
	s := d.EventTypes[typ]
	for _, req := range s.RequiredAttrs {
		if attrs[req] == "" {
			return fmt.Errorf("%w: missing attr %q for type %q", ErrInvariantViolation, req, typ)
		}
	}
	return nil
}

// ProjectTheta applica ΠΘ: riporta θ_raw nel politopo Θ^(v) (capitolo 7).
func (d Document) ProjectTheta(thetaRaw ThetaRaw) ThetaEffective {
	return ThetaEffective{
		SalienceGate: clamp(thetaRaw.SalienceGate, d.ThetaBounds.SalienceGateMin, d.ThetaBounds.SalienceGateMax),
		Eta:          clamp(thetaRaw.Eta, d.ThetaBounds.EtaMin, d.ThetaBounds.EtaMax),
	}
}

// ThetaRaw parametri «grezzi» prima della proiezione costituzionale.
type ThetaRaw struct {
	SalienceGate float64 `json:"salience_gate"`
	Eta          float64 `json:"eta"`
}

// ThetaEffective è θ_eff dopo ΠΘ — ciò che gli executor devono usare.
type ThetaEffective struct {
	SalienceGate float64 `json:"salience_gate"`
	Eta          float64 `json:"eta"`
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
