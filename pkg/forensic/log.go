// Package forensic piano append‑only separato dalla memoria comportamentale (capitoli 7, 15).
package forensic

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

// Record è byte sequence immutabile dopo scrittura.
type Record struct {
	Payload   []byte
	PrevHash  string
	SelfHash  string
}

// Log accumula record con catena di hash leggera (non blockchain hype — integrità locale).
type Log struct {
	mu       sync.Mutex
	Records  []Record
	prevHash string
}

// NewLog crea log vuoto con genesi.
func NewLog() *Log {
	return &Log{prevHash: "genesis"}
}

// Append aggiunge payload e aggiorna hash a catena.
func (l *Log) Append(payload []byte) Record {
	l.mu.Lock()
	defer l.mu.Unlock()
	sum := sha256.Sum256(append([]byte(l.prevHash+":"), payload...))
	self := hex.EncodeToString(sum[:])
	rec := Record{Payload: append([]byte(nil), payload...), PrevHash: l.prevHash, SelfHash: self}
	l.Records = append(l.Records, rec)
	l.prevHash = self
	return rec
}
