package clock

import (
	"time"

	"github.com/lacsar712/slagquen/internal/model"
)

type QuenchWindow struct {
	clk      Clock
	duration time.Duration
}

func NewQuenchWindow(clk Clock, duration time.Duration) *QuenchWindow {
	if duration <= 0 {
		duration = 2 * time.Minute
	}
	return &QuenchWindow{clk: clk, duration: duration}
}

func (w *QuenchWindow) Active(anchor time.Time) bool {
	return time.Since(anchor) < w.duration
}

func (w *QuenchWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrGranHold
}
