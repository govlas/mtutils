package mtutils

import (
	"sync"
	"sync/atomic"
)

// template type Stack(A)

var EmptyStack interface{}

type Stack interface {
	Count() int
	Push(interface{})
	Pop() interface{}
}

type nodeStack struct {
	data interface{}
	next *nodeStack
}

type stackStack struct {
	head  *nodeStack
	count int32
}

func NewStack(threadSafe bool) Stack {
	if threadSafe {
		return &tsStack{}
	} else {
		return &noTsStack{}
	}
}

func (s *stackStack) Count() int {
	return int(atomic.LoadInt32(&s.count))
}

func (s *stackStack) push(item interface{}) {
	atomic.AddInt32(&s.count, 1)

	n := &nodeStack{data: item}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
}

func (s *stackStack) pop() interface{} {
	var n *nodeStack
	if s.head != nil {
		n = s.head
		s.head = n.next
	}
	if n == nil {
		return EmptyStack
	}
	atomic.AddInt32(&s.count, -1)
	return n.data
}

type tsStack struct {
	sync.Mutex
	stackStack
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
	stackStack
}

func (s *noTsStack) Push(item interface{}) {
	s.push(item)
}

func (s *noTsStack) Pop() interface{} {
	return s.pop()
}
