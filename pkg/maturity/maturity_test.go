package maturity

import "testing"

func TestScore(t *testing.T) {
	s := Score(DefaultWeights(), Flags{KMeasured: true, ReplayCI: true})
	if s != 2 {
		t.Fatalf("got %v", s)
	}
}
