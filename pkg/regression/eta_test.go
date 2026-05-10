package regression

import "testing"

func TestStepEta(t *testing.T) {
	n := StepEta(1, 0.5, 0.2, 2)
	if n <= 1 {
		t.Fatalf("expected growth from g, got %v", n)
	}
}
