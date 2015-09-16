package mtutils

import "time"

type WaiterTask struct {
	f          func()
	tmout      time.Duration
	stopCh     chan int
	periodical bool
}

func (wt *WaiterTask) start(s <-chan int) {
	select {
	case <-time.After(wt.tmout):
		wt.f()
		if wt.periodical {
			go wt.start(s)
		}
	case <-wt.stopCh:
	case <-s:
	}
}

func (wt *WaiterTask) Stop() {
	close(wt.stopCh)
}

type Waiter struct {
	tmout   time.Duration
	taskCh  chan *WaiterTask
	cCh     chan int
	stopped bool
}

func NewWaiter(tmout time.Duration) *Waiter {
	w := &Waiter{
		tmout,
		make(chan *WaiterTask),
		make(chan int),
		false,
	}

	go w.loop()

	return w
}

func (w *Waiter) loop() {
	for task := range w.taskCh {
		go task.start(w.cCh)
	}
}

func (w *Waiter) Stop() {
	w.stopped = true
	close(w.taskCh)
	close(w.cCh)
}

func (w *Waiter) Wait(periodical bool, f func()) *WaiterTask {
	if w.stopped {
		panic("Waiter was stopped")
	}
	wt := &WaiterTask{
		f,
		w.tmout,
		make(chan int),
		periodical,
	}
	w.taskCh <- wt
	return wt
}

func WaitOnce(tm time.Duration, f func()) *WaiterTask {
	w := NewWaiter(tm)
	return w.Wait(false, func() { f(); w.Stop() })
}
