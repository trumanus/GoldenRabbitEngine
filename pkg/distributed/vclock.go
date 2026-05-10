// Package distributed modella ordinamento causale e merge (capitolo 12, Volume 2).
package distributed

// VectorClock ordinamento parziale «accade prima» tra nodi.
type VectorClock map[string]uint64

// HappensBefore verifica ordinamento parziale classico: ∀i a_i≤b_i e ∃i a_i<b_i.
func HappensBefore(a, b VectorClock) bool {
	keys := map[string]struct{}{}
	for k := range a {
		keys[k] = struct{}{}
	}
	for k := range b {
		keys[k] = struct{}{}
	}
	strict := false
	for k := range keys {
		va := a[k]
		vb := b[k]
		if va > vb {
			return false
		}
		if va < vb {
			strict = true
		}
	}
	return strict
}

// MergeComponentWise massimo componente per stato derivato (merge 𝒟𝓂 𝓁 scoped).
func MergeComponentWise(shards ...VectorClock) VectorClock {
	out := VectorClock{}
	for _, s := range shards {
		for k, v := range s {
			if v > out[k] {
				out[k] = v
			}
		}
	}
	return out
}
