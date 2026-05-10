// Package outbox pattern transazione + pubblicazione messaggio (capitoli 2 e 15, Volume 2).
package outbox

import (
	"errors"
	"sync"
)

var ErrPublishFailed = errors.New("outbox: publish failed")

// Message da inviare sul bus tipizzato (payload opaco).
type Message struct {
	ID      string
	Topic   string
	Payload []byte
}

// Handler pubblica verso transport reale; in laboratorio no‑op o slice collector.
type Handler func(m Message) error

// Store mantiene outbox in memoria con mutex.
type Store struct {
	mu       sync.Mutex
	Messages []Message
	publish  Handler
}

// New crea store con handler (può essere nil → accumulate only).
func New(publish Handler) *Store {
	return &Store{publish: publish}
}

// StageWithinTx esegue fn e solo se ok appende messaggio — pattern «stesso commit logico».
func (s *Store) StageWithinTx(msg Message, fn func() error) error {
	if fn == nil {
		return errors.New("outbox: nil fn")
	}
	if err := fn(); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.publish != nil {
		if err := s.publish(msg); err != nil {
			return ErrPublishFailed
		}
	}
	s.Messages = append(s.Messages, msg)
	return nil
}
