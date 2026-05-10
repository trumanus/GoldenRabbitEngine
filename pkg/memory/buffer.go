package memory

import (
	"encoding/json"
)

// ShortBuffer è M_B: finestra limitata di episodi non ancora promossi (capitolo 8).
type ShortBuffer struct {
	Max   int
	Items [][]byte `json:"items"`
}

// Push aggiunge payload serializzato e taglia FIFO.
func (b *ShortBuffer) Push(payload any) error {
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	b.Items = append(b.Items, raw)
	if b.Max > 0 && len(b.Items) > b.Max {
		b.Items = b.Items[len(b.Items)-b.Max:]
	}
	return nil
}
