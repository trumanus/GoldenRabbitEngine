package distributed

import "testing"

func TestHappensBefore(t *testing.T) {
	a := VectorClock{"s1": 1, "s2": 0}
	b := VectorClock{"s1": 2, "s2": 1}
	if !HappensBefore(a, b) {
		t.Fatal("expected a before b")
	}
	if HappensBefore(b, a) {
		t.Fatal("not concurrent")
	}
}
