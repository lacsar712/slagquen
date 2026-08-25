package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lacsar712/slagquen/internal/config"
	"github.com/lacsar712/slagquen/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	anchor := a.clk.Now().Add(-3 * time.Minute)
	err = a.ConfirmGranHold(context.Background(), anchor)
	if err == nil {
		t.Fatal("expected gradient hold error")
	}
	if !errors.Is(err, model.ErrGranHold) {
		t.Fatalf("expected ErrGranHold, got %v", err)
	}
	_ = time.Second
}
