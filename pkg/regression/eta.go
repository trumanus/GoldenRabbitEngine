// Package regression aggiorna η secondo η(t+1)=η(t)+α·g(𝒟𝓂𝓁,S)-β·η(t) (capitolo 14, Volume 2).
package regression

// StepEta applica un passo discreto con contributo g osservato (curvatura dal passato promosso).
func StepEta(eta, alpha, beta, g float64) float64 {
	return eta + alpha*g - beta*eta
}
