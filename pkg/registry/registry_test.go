package registry

import "testing"

func TestHashRegistryStable(t *testing.T) {
	d := DefaultDocument()
	h1, err := HashRegistry(d)
	if err != nil {
		t.Fatal(err)
	}
	h2, err := HashRegistry(d)
	if err != nil {
		t.Fatal(err)
	}
	if h1 != h2 {
		t.Fatalf("hash unstable: %s vs %s", h1, h2)
	}
}

func TestProjectTheta(t *testing.T) {
	d := DefaultDocument()
	d.ThetaBounds.SalienceGateMin = 0.2
	d.ThetaBounds.SalienceGateMax = 0.8
	d.ThetaBounds.EtaMin = 0
	d.ThetaBounds.EtaMax = 0.5

	e := d.ProjectTheta(ThetaRaw{SalienceGate: 999, Eta: -9})
	if e.SalienceGate != 0.8 || e.Eta != 0 {
		t.Fatalf("unexpected projection: %+v", e)
	}
}

func TestValidateAttrs(t *testing.T) {
	d := DefaultDocument()
	err := d.ValidateAttrs("order.confirmed", map[string]string{"tenant": "t1"})
	if err == nil {
		t.Fatal("expected missing order_id error")
	}
	err = d.ValidateAttrs("order.confirmed", map[string]string{"tenant": "t1", "order_id": "o1"})
	if err != nil {
		t.Fatal(err)
	}
}
