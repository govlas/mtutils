package mtutils

import "time"

type Waiter chan func()

func NewWaiter(tmout time.Duration) Waiter {
	w := make(Waiter)

	go func() {
		for f := range w {
			go WaitOnce(tmout, f)
		}
	}()

	return w
}

func (w Waiter) Stop() {
	close(w)
}

func (w Waiter) Wait(periodical bool, f func()) {
	if periodical {
		w <- func() {
			f()
			w.Wait(periodical, f)
		}
	} else {
		w <- f
	}
}

func WaitOnce(tm time.Duration, f func()) {
	time.Sleep(tm)
	f()
}
