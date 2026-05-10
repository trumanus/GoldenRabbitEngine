package budget

import "testing"

func TestSpendAndRecharge(t *testing.T) {
	l, err := NewLedger([]float64{10, 10, 10})
	if err != nil {
		t.Fatal(err)
	}
	if err := l.Spend([]float64{1, 2, 3}); err != nil {
		t.Fatal(err)
	}
	if l.K[2] != 7 {
		t.Fatalf("promotion channel: %v", l.K)
	}
	if err := l.EndTick(); err != nil {
		t.Fatal(err)
	}
	if l.K[2] <= 7 {
		t.Fatalf("expected recharge, got %v", l.K)
	}
}

func TestFeasible(t *testing.T) {
	if !Feasible([]float64{5, 5}, []float64{3, 5}) {
		t.Fatal("should be feasible")
	}
	if Feasible([]float64{1}, []float64{2}) {
		t.Fatal("should not")
	}
}
