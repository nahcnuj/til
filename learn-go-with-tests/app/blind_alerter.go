package app

import (
	"fmt"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// idiomatically expose also the func type for an interface that has just one function
type BlindAlerterFunc func(duration time.Duration, amount int)

func (f BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	f(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Printf("Blind is now %d\n", amount)
	})
}
