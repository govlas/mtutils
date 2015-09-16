package mtutils_test

import (
	"sync"
	"testing"

	"github.com/govlas/mtutils"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := mtutils.NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, s.Pop(), 3, "TestStack")
	assert.Equal(t, s.Pop(), 2, "TestStack")
	assert.Equal(t, s.Pop(), 1, "TestStack")

	assert.Equal(t, s.Pop(), nil, "TestStack")
}

func push(v int, s *mtutils.Stack, wg *sync.WaitGroup) {
	defer wg.Done()
	s.Push(v)
}
func TestStackCount(t *testing.T) {

	s := mtutils.NewStack()
	var wg sync.WaitGroup
	wg.Add(3)
	go push(1, s, &wg)
	go push(2, s, &wg)
	go push(3, s, &wg)

	wg.Wait()

	assert.Equal(t, s.Count(), 3, "TestStackCount")

	for s.Count() > 0 {
		s.Pop()
	}
	assert.Equal(t, s.Count(), 0, "TestStackCount")
}
