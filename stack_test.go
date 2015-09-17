package mtutils_test

import (
	"testing"

	"github.com/govlas/mtutils"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	s := mtutils.NewStack(false)
	const cnt = 100
	for i := 0; i < cnt; i++ {
		s.Push(i)
	}

	assert.Equal(t, s.Count(), cnt, "TestStack")

	a := cnt - 1
	for s.Count() > 0 {
		b := s.Pop()
		if !assert.Equal(t, b, a, "TestStack") {
			break
		}
		a--
	}
	assert.Equal(t, s.Count(), 0, "TestStack")
	assert.Equal(t, s.Pop(), nil, "TestStack")
}

func TestStackMt(t *testing.T) {
	s := mtutils.NewStack(true)
	ch := make(chan int)
	wait := make(chan int)

	go func() {
		for i := range ch {
			s.Push(i)
		}
	}()

	go func() {
		for {
			b := s.Pop()

			if b == 0 {
				wait <- 1
				break
			}
		}
	}()

	const cnt = 100
	for i := 0; i < cnt; i++ {
		ch <- i
	}
	close(ch)
	<-wait
}
