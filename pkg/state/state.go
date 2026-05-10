// Package state modella uno snapshot ridotto di X(t) (capitolo 14, Volume 2).
package state

import (
	"encoding/json"

	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
)

// Snapshot è una vista serializzabile di componenti osservabili per replay e runbook.
type Snapshot struct {
	RegistryVersion string `json:"registry_version"`
	RegistryHash    string `json:"registry_hash"`

	SessionSummary string  `json:"session_summary,omitempty"` // S — riassunto carry-over
	KBudget        float64 `json:"k_budget"`                  // K scalare semplificato
	Eta            float64 `json:"eta"`                       // η regressivo

	ThetaRaw      registry.ThetaRaw      `json:"theta_raw"`
	ThetaEff      registry.ThetaEffective `json:"theta_eff"`

	LongMemory json.RawMessage `json:"long_memory"` // delegato al pkg memory per forma canonica
}
