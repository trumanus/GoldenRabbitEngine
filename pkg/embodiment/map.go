// Package embodiment mappa z_phys → contributi di stato (capitolo 11, Volume 2).
package embodiment

// Features normalizza letture fisiche in chiavi numeriche osservabili.
type Features map[string]float64

// Encode applica scala e clamp dichiarati (senza dipendenze robotiche hardcoded).
func Encode(in Features, scale map[string]float64) Features {
	out := make(Features, len(in))
	for k, v := range in {
		s := scale[k]
		if s == 0 {
			s = 1
		}
		out[k] = v / s
	}
	return out
}
