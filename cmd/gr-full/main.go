// Command gr-full dimostra kernel: budget, omeostasi, λ, forensic, outbox, maturity (Volume 2).
//
// Codice creato da Andrea Lagomarsini (trumanus) <trumanus@gmail.com>.
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/trumanus/GoldenRabbitEngine/pkg/act"
	"github.com/trumanus/GoldenRabbitEngine/pkg/budget"
	"github.com/trumanus/GoldenRabbitEngine/pkg/conflict"
	"github.com/trumanus/GoldenRabbitEngine/pkg/distributed"
	"github.com/trumanus/GoldenRabbitEngine/pkg/embodiment"
	"github.com/trumanus/GoldenRabbitEngine/pkg/identity"
	"github.com/trumanus/GoldenRabbitEngine/pkg/kernel"
	"github.com/trumanus/GoldenRabbitEngine/pkg/maturity"
	"github.com/trumanus/GoldenRabbitEngine/pkg/memory"
	"github.com/trumanus/GoldenRabbitEngine/pkg/outbox"
	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
	"github.com/trumanus/GoldenRabbitEngine/pkg/security"
	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
	"github.com/trumanus/GoldenRabbitEngine/pkg/values"
)

func main() {
	log.SetFlags(0)
	reg := registry.DefaultDocument()
	h, err := registry.HashRegistry(reg)
	if err != nil {
		log.Fatal(err)
	}
	c := state.NewCognitive(reg, h, 0, registry.ThetaRaw{SalienceGate: 0.5, Eta: 0.1})
	led, err := budget.NewLedger([]float64{20, 20, 20})
	if err != nil {
		log.Fatal(err)
	}
	sys := kernel.NewSystem(reg, h, c, led)
	_ = sys.InitLambda([]float64{0.2}, []float64{0.8}, 0.1)
	sys.SyncPromotionBudget()
	sys.Tau = values.Tensions{"latency": 0.5, "hr": 2}
	sys.Homeo.Density = 0.7

	sys.Forensic.Append([]byte(`{"kind":"boot","registry":"` + h[:12] + `"}`))

	enc := embodiment.Encode(embodiment.Features{"temp_c": 42}, map[string]float64{"temp_c": 100})
	fmt.Println("=== gr-full (kernel demo) ===")
	fmt.Println("embodiment z_phys:", enc)

	if err := security.TenantScope("t1", "t1", true); err != nil {
		log.Fatal(err)
	}
	if err := security.TenantScope("t1", "t2", true); err == nil {
		log.Fatal("expected tenant leak error")
	}

	vc1 := distributed.VectorClock{"a": 1}
	vc2 := distributed.VectorClock{"a": 2, "b": 1}
	fmt.Println("vector clock a→b:", distributed.HappensBefore(vc1, vc2))

	res, ok := conflict.Resolve(sys.Tau, []conflict.Candidate{
		{Name: "rush", Loss: 0, Risk: 3, Tags: []string{"hr"}},
		{Name: "careful", Loss: 2, Risk: 0.2, Tags: []string{"hr"}},
	})
	if ok {
		fmt.Println("conflict:", res.Chosen.Name, res.Score)
	}

	a, ok := act.Select(led, []act.Action{
		{Name: "deep", Cost: []float64{8, 8, 8}},
		{Name: "cheap", Cost: []float64{1, 1, 2}},
	}, []float64{0, 5})
	if ok {
		fmt.Println("ACT chose:", a.Name)
	}

	msg := outbox.Message{ID: "demo-1", Topic: "goldenrabbit.events", Payload: []byte(`{}`)}
	if err := sys.Outbox.StageWithinTx(msg, func() error {
		return sys.Multi.Scene.Push(map[string]string{"session": "ping"})
	}); err != nil {
		log.Fatal(err)
	}

	if err := sys.Tick(12 /* E osservato */, []float64{1, 1, 2}); err != nil {
		log.Fatal(err)
	}

	mScore := maturity.Score(maturity.DefaultWeights(), maturity.Flags{
		KMeasured: true, PiSigned: false, ReplayCI: true, TenantIsolated: true,
	})
	fmt.Println("maturity M:", mScore)

	phi := identity.FromState(*sys.Cognitive)
	fmt.Println("identity:", phi.String())

	item := memory.CompositeScore(0.9, 0.2, 0.5, 0.1, memory.RetrievalWeights{AlphaSim: 1, BetaRisk: 2, GammaAge: 0.5, DeltaViol: 3})
	fmt.Printf("retrieval score: total=%.4f\n", item.Total)

	out, _ := json.MarshalIndent(sys.Multi.MassByLayer(), "", "  ")
	fmt.Println("mass by layer:", string(out))
	fmt.Println("forensic records:", len(sys.Forensic.Records))
	fmt.Println("outbox messages:", len(sys.Outbox.Messages))
}
