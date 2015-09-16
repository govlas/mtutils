package mtutils

import (
	"sync"
	"sync/atomic"
)

type node struct {
	data interface{}
	next *node
}

type Stack struct {
	sync.Mutex
	head  *node
	count int32
}

func NewStack() *Stack {
	s := new(Stack)
	return s
}

func (s *Stack) Count() int {
	return int(atomic.LoadInt32(&s.count))
}

func (s *Stack) Push(item interface{}) {
	s.Lock()
	defer s.Unlock()

	atomic.AddInt32(&s.count, 1)

	n := &node{data: item}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
}

func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()

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
