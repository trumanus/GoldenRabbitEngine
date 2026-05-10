package outbox

import (
	"errors"
	"testing"
)

func TestStageRollback(t *testing.T) {
	s := New(nil)
	err := s.StageWithinTx(Message{ID: "1"}, func() error { return errors.New("fail") })
	if err == nil {
		t.Fatal("expected error")
	}
	if len(s.Messages) != 0 {
		t.Fatal("should not append")
	}
}

func TestStageCommit(t *testing.T) {
	s := New(nil)
	if err := s.StageWithinTx(Message{ID: "1"}, func() error { return nil }); err != nil {
		t.Fatal(err)
	}
	if len(s.Messages) != 1 {
		t.Fatal("expected append")
	}
}
