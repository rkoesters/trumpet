package trumpet

import (
	"time"
)

// A Scheduler learns from calls to Train() and decides when by sending
// a signal through Chan().
type Scheduler interface {
	Train(d time.Time)
	Chan() <-chan struct{}
}
