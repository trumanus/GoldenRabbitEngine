package memory

// Layer stratifica ℓ per 𝒟𝓂 𝓁 e ranking (capitolo 8, Volume 2).
type Layer string

const (
	LayerScene      Layer = "scene"      // M_B — scena corta
	LayerEpisodic   Layer = "episodic"
	LayerSemantic   Layer = "semantic"
	LayerCognitive  Layer = "cognitive" // piega P_ACT / η
	LayerRelational Layer = "relational"
	LayerForensic   Layer = "forensic" // append‑only separato nel pkg forensic
)
