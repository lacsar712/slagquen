package app

import (
	"context"
	"errors"
	"testing"

	"github.com/lacsar712/slagquen/internal/config"
	"github.com/lacsar712/slagquen/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	err = a.HandleQuenchTrip(context.Background(), model.TowerID(a.cfg.TowerID), 95.0)
	if err == nil {
		t.Fatal("expected heat overtemperature error")
	}
	if !errors.Is(err, model.ErrQuenchTrip) {
		t.Fatalf("expected ErrQuenchTrip, got %v", err)
	}
}
