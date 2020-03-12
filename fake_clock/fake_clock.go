package fake_clock

import (
	"time"
)

// FakeClock is used in testing to simulate time flowing
//
// Each Tick() call increment elapsed time by 1 second
// starting from 2010-01-01T00:00:00Z
type FakeClock struct {
	start time.Time
}

//New creates a new clock initialized at 2010-01-01
func New() *FakeClock {
	t, _ := time.Parse(time.RFC3339, "2010-01-01T00:00:00Z")
	return &FakeClock{
		start: t,
	}
}

func (c *FakeClock) Tick() {
	c.start = c.start.Add(time.Second)
}

func (c *FakeClock) Now() time.Time {
	return c.start
}

func (c *FakeClock) RFC3339() string {
	return c.start.Format(time.RFC3339Nano)
}
