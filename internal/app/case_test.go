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
	err = a.ValidateSlagDrift(context.Background(), 25.0)
	if err == nil {
		t.Fatal("expected moisture drift violation")
	}
	if !errors.Is(err, model.ErrSlagDrift) {
		t.Fatalf("expected ErrSlagDrift, got %v", err)
	}
}
