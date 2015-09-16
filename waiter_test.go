package mtutils_test

import (
	"testing"
	"time"

	"github.com/govlas/mtutils"
	"github.com/stretchr/testify/assert"
)

func TestOnceWaiter(t *testing.T) {
	inCh := make(chan time.Time)
	now := time.Now()
	mtutils.WaitOnce(time.Second, func() { inCh <- time.Now() })

	tm := <-inCh
	dur := tm.Sub(now)
	t.Log(dur)
	assert.Equal(t, dur/time.Second*time.Second, time.Second, "TestOnceWaiter")
}

func TestPeriodicalWaiter(t *testing.T) {

	inCh := make(chan int)
	now := time.Now()
	w := mtutils.NewWaiter(time.Second)
	w.Wait(true, func() { inCh <- 1 })

	var i int
	for _ = range inCh {
		dur := time.Now().Sub(now)
		t.Log(i, dur)
		assert.Equal(t, dur/time.Second*time.Second, time.Second, "TestPeriodicalWaiter")
		now = time.Now()
		i++
		if i >= 3 {
			w.Stop()
			close(inCh)
		}
	}
}

func TestTaskWaiterStop(t *testing.T) {
	inCh := make(chan int)
	wt := mtutils.WaitOnce(time.Second, func() { inCh <- 1 })

	wt.Stop()

	select {
	case <-inCh:
		t.Fatal("WaiterTask.Stop() does not working")
	case <-time.After(time.Second * 2):
	}

}
