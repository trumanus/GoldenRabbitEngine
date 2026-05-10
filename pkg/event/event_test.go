package event

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/trumanus/GoldenRabbitEngine/pkg/registry"
)

func TestPipelineBotSalience(t *testing.T) {
	p := Pipeline{
		Reg: registry.DefaultDocument(),
		Enc: FunctionalEncoder(func(in RawInput) (Encoded, error) {
			return Encoded{
				RequestID:   in.RequestID,
				Tenant:      in.Tenant,
				ProjectTick: in.DeclaredTick,
			}, nil
		}),
	}
	in := RawInput{
		RequestID:    "r1",
		ReceivedAt:   time.Now(),
		Tenant:       "t",
		Payload:      json.RawMessage(`{}`),
		DeclaredTick: 1,
	}
	_, err := p.FromRaw(in, "order.confirmed", map[string]string{"tenant": "t", "order_id": "o"}, 0.1)
	if err != ErrBot {
		t.Fatalf("expected ErrBot, got %v", err)
	}
}

func TestPipelineTyped(t *testing.T) {
	p := Pipeline{
		Reg: registry.DefaultDocument(),
		Enc: FunctionalEncoder(func(in RawInput) (Encoded, error) {
			return Encoded{RequestID: in.RequestID, Tenant: in.Tenant, ProjectTick: 7, PayloadRef: "sha256:abc"}, nil
		}),
	}
	in := RawInput{
		RequestID:    "r2",
		Tenant:       "t",
		Payload:      json.RawMessage(`{}`),
		Idempotency:  "intent-1",
		Correlation:  "corr-9",
		DeclaredTick: 2,
	}
	e, err := p.FromRaw(in, "order.confirmed", map[string]string{"tenant": "t", "order_id": "o42"}, 0.99)
	if err != nil {
		t.Fatal(err)
	}
	if e.IdempotencyKey != "intent-1" || e.ProjectTick != 7 || e.PayloadRef != "sha256:abc" {
		t.Fatalf("unexpected typed event: %+v", e)
	}
}
