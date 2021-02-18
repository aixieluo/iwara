package untils

import (
	"github.com/robfig/cron/v3"
)

const timing string  = "0 0 0 1/1 * ?"

func Schedules(fn func()) {
	c := cron.New(cron.WithSeconds())

	_, _ = c.AddFunc(timing, func() {
		fn()
	})

	go c.Start()
}