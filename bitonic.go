package mtutils

import (
	"sort"
	"sync"
)

const (
	SORT_ASC  bool = true
	SORT_DESC bool = false
)

// BitonicSort sorts collection of elements by bitonic sort algorithm.
// Based on article: http://farazdagi.com/blog/2015/bitonic-sort-in-go/
func BitonicSort(s sort.Interface) {
	bitonicSort(s, SORT_ASC)
}

// BitonicReverse sorts collection of elements by bitonic sort algorithm in reverse direction.
// Based on article: http://farazdagi.com/blog/2015/bitonic-sort-in-go/
func BitonicReverse(s sort.Interface) {
	bitonicSort(s, SORT_DESC)
}

func bitonicSort(s sort.Interface, dir bool) {
	(&arr{Interface: s}).startSorting(dir)
}

type arr struct {
	sort.Interface
}

func (a *arr) startSorting(dir bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	go a.spSort(0, a.Len(), dir, &wg)
	wg.Wait()
}

func (a *arr) spSort(lo int, n int, dir bool, wg *sync.WaitGroup) {
	defer wg.Done()
	if n > 1 {
		m := n / 2
		var iwg sync.WaitGroup

		iwg.Add(2)

		go a.spSort(lo, m, !dir, &iwg)
		go a.spSort(lo+m, n-m, dir, &iwg)

		iwg.Wait()

		a.spMerge(lo, n, dir)
	}
}

func (a *arr) spMerge(lo int, n int, dir bool) {
	if n > 1 {

		m := 1
		for m < n {
			m = m << 1
		}
		m = m >> 1

		for i := lo; i < lo+n-m; i++ {
			if a.Less(i, i+m) != dir {
				a.Swap(i, i+m)
			}
		}

		a.spMerge(lo, m, dir)
		a.spMerge(lo+m, n-m, dir)

	}

}
