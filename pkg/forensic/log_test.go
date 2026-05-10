package forensic

import "testing"

func TestChain(t *testing.T) {
	l := NewLog()
	a := l.Append([]byte("e1"))
	b := l.Append([]byte("e2"))
	if b.PrevHash != a.SelfHash {
		t.Fatalf("broken chain")
	}
}
