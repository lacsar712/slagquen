package app

import (
	"context"
	"fmt"
	"time"

	"github.com/lacsar712/slagquen/internal/model"
)

func (a *App) ValidateSlagDrift(ctx context.Context, moistPct float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	limit := a.cfg.TargetMoistPct + a.cfg.MaxGradientDeltaPct
	if moistPct <= limit {
		return nil
	}
	return fmt.Errorf("moisture: %w", model.ErrSlagDrift)
}

func (a *App) ConfirmGranHold(ctx context.Context, anchor time.Time) error {
	if a.avgWindow == nil {
		return model.Wrap("app", "window", model.ErrGranHold)
	}
	if err := a.avgWindow.Require(anchor); err != nil {
		return fmt.Errorf("gradient hold: window not satisfied")
	}
	return nil
}
