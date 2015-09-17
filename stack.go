package mtutils

import (
	"sync"
	"sync/atomic"
)

type Stack interface {
	Count() int
	Push(interface{})
	Pop() interface{}
}

type node struct {
	data interface{}
	next *node
}

type stack struct {
	head  *node
	count int32
}

func NewStack(threadSafe bool) Stack {
	if threadSafe {
		return &tsStack{}
	} else {
		return &noTsStack{}
	}
}

func (s *stack) Count() int {
	return int(atomic.LoadInt32(&s.count))
}

func (s *stack) push(item interface{}) {
	atomic.AddInt32(&s.count, 1)

	n := &node{data: item}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
}

func (s *stack) pop() interface{} {
	var n *node
	if s.head != nil {
		n = s.head
		s.head = n.next
	}
	if n == nil {
		return nil
	}
	atomic.AddInt32(&s.count, -1)
	return n.data
}

type tsStack struct {
	sync.Mutex
	stack
}

func (s *tsStack) Push(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.push(item)
}

func (s *tsStack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()
	return s.pop()
}

type noTsStack struct {
	stack
}

func (s *noTsStack) Push(item interface{}) {
	s.push(item)
}

func (s *noTsStack) Pop() interface{} {
	return s.pop()
}
