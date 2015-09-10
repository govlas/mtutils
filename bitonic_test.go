package mtutils_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/govlas/mtutils"
	"github.com/stretchr/testify/assert"
)

func TestBitonicSort(t *testing.T) {

	const sliceSize = 101

	src := makeTempSlice(sliceSize)
	src1 := sort.IntSlice(make([]int, len(src)))
	copy(src1, src)
	mtutils.BitonicSort(src1)
	sort.Sort(src)
	assert.Equal(t, src, src1, "BitonicSort")

	src = makeTempSlice(sliceSize)
	copy(src1, src)
	mtutils.BitonicReverse(sort.IntSlice(src1))
	sort.Sort(sort.Reverse(src))
	assert.Equal(t, src, src1, "BitonicReverse")

}

func makeTempSlice(n int) sort.IntSlice {
	rand.Seed(time.Now().UnixNano())
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = rand.Intn(100)
	}
	return sort.IntSlice(ret)
}

func BenchmarkBitonicSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		src := makeTempSlice(1 << 16)
		b.StartTimer()
		mtutils.BitonicSort(src)
		b.StopTimer()
	}
}

func BenchmarkStandartSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		src := makeTempSlice(1 << 16)
		b.StartTimer()
		sort.Sort(src)
		b.StopTimer()
	}
}
