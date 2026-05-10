// Command gr-demo mostra catena registry → evento → promozione → snapshot (Volume 2).
//
// Codice creato da Andrea Lagomarsini (trumanus) <trumanus@gmail.com>.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/trumanus/GoldenRabbitEngine/pkg/event"
	"github.com/trumanus/GoldenRabbitEngine/pkg/promotion"
	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
	"github.com/trumanus/GoldenRabbitEngine/pkg/state"
)

func main() {
	log.SetFlags(0)
	reg := registry.DefaultDocument()
	h, err := registry.HashRegistry(reg)
	if err != nil {
		log.Fatal(err)
	}

	raw := registry.ThetaRaw{SalienceGate: 0.9, Eta: 0.12}
	thetaEff := reg.ProjectTheta(raw)
	fmt.Println("=== Golden Rabbit engine (demo) ===")
	fmt.Printf("Registry version: %s\nHash Registry:   %s\nθ_eff:           salience_gate=%.3f η=%.3f\n\n",
		reg.Version, h, thetaEff.SalienceGate, thetaEff.Eta)

	pipe := event.Pipeline{
		Reg: reg,
		Enc: event.FunctionalEncoder(func(in event.RawInput) (event.Encoded, error) {
			return event.Encoded{
				RequestID:   in.RequestID,
				Tenant:      in.Tenant,
				ProjectTick: in.DeclaredTick,
				PayloadRef:  "sha256:demo",
			}, nil
		}),
	}

	rawIn := event.RawInput{
		RequestID:    "req-demo-1",
		ReceivedAt:   time.Now().UTC(),
		Tenant:       "tenant-acme",
		Payload:      []byte(`{"sku":"R42"}`),
		Idempotency:  "idem-demo-1",
		Correlation:  "corr-demo-1",
		DeclaredTick: 42,
	}

	sigma := 0.8
	ev, err := pipe.FromRaw(rawIn, "preference.promoted", map[string]string{
		"tenant":        "tenant-acme",
		"scope":         "procurement",
		"rationale_ref": "doc://postmortem/2026-01",
	}, sigma)
	if err != nil {
		log.Fatal(err)
	}

	c := state.NewCognitive(reg, h, 50, raw)
	eng := promotion.NewEngine()
	intent := promotion.Intent{
		IdempotencyKey:       ev.IdempotencyKey,
		ThetaVersion:         "demo-0.1.0",
		RegistryHashExpected: h,
		KappaPi:              5,
	}

	applied, err := eng.Apply(&c, intent, func() error {
		return promotion.AppendTraceFromEvent(&c, ev, "trace-demo-1", 3, "preference")
	})
	if err != nil {
		log.Fatal(err)
	}

	snap, err := c.Snapshot()
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	fmt.Printf("Event typed: %s (tick=%d)\nPromotion applied: %v\n\nSnapshot JSON:\n", ev.Type, ev.ProjectTick, applied)
	if err := enc.Encode(snap); err != nil {
		log.Fatal(err)
	}
}
