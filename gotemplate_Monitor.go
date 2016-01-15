package mtutils

import (
	"runtime"
	"sync"
)

// template type Monitor(A)

type AccessFuncMonitor func(interface{})

type Monitor struct {
	sync.Mutex
	obj interface{}
}

func NewMonitor(obj interface{}) *Monitor {
	ret := new(Monitor)
	ret.obj = obj
	runtime.SetFinalizer(ret, func(m *Monitor) {
		m.Unlock()
	})
	return ret
}

func (m *Monitor) Capture() interface{} {
	m.Lock()
	return m.obj
}

func (m *Monitor) Release() {
	m.Unlock()
}

func (m *Monitor) Access(f AccessFuncMonitor) {
	m.Lock()
	defer m.Unlock()
	f(m.obj)
}
