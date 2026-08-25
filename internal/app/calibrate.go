package app

import (
	"context"
	"fmt"
	"time"

	"github.com/lacsar712/slagquen/internal/model"
)

// CalibrateProbe allows acceptance tests to inject feed calibration faults.
var CalibrateProbe func(ctx context.Context) error

const feedTempLimitC = 55.0

func (a *App) CalibrateFeed(ctx context.Context, tower model.TowerID, holder string) error {
	if err := a.feedLeases.Require(tower, holder, 30*time.Second); err != nil {
		return err
	}
	// Release on every return path: a probe fault (e.g. disconnected mid-zero)
	// must not leave the feed lease held, or the tower stays "occupied" and
	// downstream quench/spray acquisition is blocked until a full restart.
	defer a.feedLeases.ReleaseHolder(tower, holder)
	if CalibrateProbe != nil {
		if err := CalibrateProbe(ctx); err != nil {
			return fmt.Errorf("calibrate: %w", err)
		}
	}
	return nil
}
