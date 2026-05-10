package state

import (
	"github.com/trumanus/GoldenRabbitEngine/pkg/memory"
	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
)

// Cognitive tiene lo stato mutabile minimo per demo e test di replay (capitolo 14: X).
type Cognitive struct {
	Reg      registry.Document
	RegHash  string
	Mem      memory.LongMemory
	K        float64
	ThetaRaw registry.ThetaRaw
	ThetaEff registry.ThetaEffective
	Eta      float64
	Tick     int64
}

// NewCognitive costruisce uno stato con η e θ proiettati.
func NewCognitive(reg registry.Document, regHash string, k0 float64, raw registry.ThetaRaw) Cognitive {
	c := Cognitive{
		Reg:      reg,
		RegHash:  regHash,
		K:        k0,
		ThetaRaw: raw,
	}
	c.ThetaEff = reg.ProjectTheta(raw)
	c.Eta = c.ThetaEff.Eta // η osservabile allineato a ΠΘ (capitolo 7)
	return c
}

// RefreshTheta ricalcola θ_eff dopo mutazione costituzionale o raw.
func (c *Cognitive) RefreshTheta() {
	c.ThetaEff = c.Reg.ProjectTheta(c.ThetaRaw)
}

// Snapshot serializza osservabili per audit (cap. 15 runbook).
func (c Cognitive) Snapshot() (Snapshot, error) {
	lm, err := c.Mem.MarshalJSONBytes()
	if err != nil {
		return Snapshot{}, err
	}
	return Snapshot{
		RegistryVersion: c.Reg.Version,
		RegistryHash:    c.RegHash,
		KBudget:         c.K,
		Eta:             c.Eta,
		ThetaRaw:        c.ThetaRaw,
		ThetaEff:        c.ThetaEff,
		LongMemory:      lm,
	}, nil
}
