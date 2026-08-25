package app

import (
	"context"
	"time"

	"github.com/lacsar712/slagquen/internal/clock"
)

type QuenchRamp struct {
	clk   clock.Clock
	tick  time.Duration
	steps int
}

func NewQuenchRamp(clk clock.Clock, tick time.Duration, steps int) *QuenchRamp {
	if steps <= 0 {
		steps = 40
	}
	return &QuenchRamp{clk: clk, tick: tick, steps: steps}
}

func (r *QuenchRamp) Ramp(ctx context.Context, target float64, apply func(float64)) error {
	step := target / float64(r.steps)
	if step <= 0 {
		step = 0.5
	}
	cur := 0.0
	for cur < target {
		cur += step
		if cur > target {
			cur = target
		}
		apply(cur)
		if pc, ok := r.clk.(*clock.ProcessClock); ok {
			pc.Step()
		}
		time.Sleep(2 * time.Millisecond)
	}
	return nil
}

func (a *App) RunQuenchRamp(ctx context.Context, target float64) error {
	return a.dryRamp.Ramp(ctx, target, func(v float64) { _ = v })
}
