package sched

import (
	"time"
)

type Scheduler struct {
	schedule []time.Time
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		schedule: make([]time.Time, 0),
	}
}

func (s *Scheduler) Train(d time.Time) {

}
