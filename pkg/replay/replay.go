// Package replay fornisce harness deterministico per ripetere promozioni (capitoli 2, 15, Volume 2).
package replay

import (
	"github.com/trumanus/GoldenRabbitEngine/pkg/promotion"
	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
)

// Step è una promozione registrata in ordine di tick di progetto.
type Step struct {
	Tick   int64
	Intent promotion.Intent
	Delta  func(*state.Cognitive) error
}

// Run applica gli step in ordine; idempotenza interna ai singoli intent è rispettata dal promotion.Engine.
func Run(c *state.Cognitive, eng *promotion.Engine, steps []Step) error {
	for _, s := range steps {
		c.Tick = s.Tick
		_, err := eng.Apply(c, s.Intent, func() error {
			if s.Delta == nil {
				return nil
			}
			return s.Delta(c)
		})
		if err != nil {
			return err
		}
	}
	return nil
}
