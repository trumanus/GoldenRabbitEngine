// Package values modella tensioni valoriali τ (capitoli 9–10, Volume 2).
package values

// Tensions è una mappa nome→intensità τ_k ≥ 0 per arbitrare trade‑off.
type Tensions map[string]float64

// Weight prende il peso di una tensione nota (0 se assente).
func (t Tensions) Weight(key string) float64 {
	if t == nil {
		return 0
	}
	return t[key]
}
