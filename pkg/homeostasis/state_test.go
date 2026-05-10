package homeostasis

import "testing"

func TestPIDMovesCommand(t *testing.T) {
	s := NewPID(10, 0.1, 0.01, 0.05)
	u1 := s.Step(12)
	u2 := s.Step(12)
	if u1 == u2 {
		t.Fatalf("PID should react to integral/derivative: %v %v", u1, u2)
	}
}
