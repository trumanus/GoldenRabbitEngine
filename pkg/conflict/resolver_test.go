package conflict

import (
	"testing"

	"github.com/trumanus/GoldenRabbitEngine/pkg/values"
)

func TestResolvePrefersLowRiskWhenTauHigh(t *testing.T) {
	tau := values.Tensions{"hr": 10}
	a := Candidate{Name: "fast", Loss: 0, Risk: 2, Tags: []string{"hr"}}
	b := Candidate{Name: "slow", Loss: 1, Risk: 0.1, Tags: []string{"hr"}}
	r, ok := Resolve(tau, []Candidate{a, b})
	if !ok || r.Chosen.Name != "slow" {
		t.Fatalf("%+v", r)
	}
}
