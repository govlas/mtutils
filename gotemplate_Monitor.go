package mtutils

import (
	"runtime"
	"sync"
)

// template type Monitor(A)

type MonitoredObj interface{}
type AccessFuncMonitor func(MonitoredObj)

type Monitor struct {
	sync.Mutex
	obj MonitoredObj
}

func NewMonitor(obj MonitoredObj) *Monitor {
	ret := new(Monitor)
	ret.obj = obj
	runtime.SetFinalizer(ret, func(m *Monitor) {
		m.Unlock()
	})
	return ret
}

func (m *Monitor) Capture() MonitoredObj {
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
