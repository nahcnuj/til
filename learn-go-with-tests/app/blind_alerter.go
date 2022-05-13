package app

import (
	"fmt"
	"io"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// idiomatically expose also the func type for an interface that has just one function
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

func (f BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	f(duration, amount, to)
}

func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
