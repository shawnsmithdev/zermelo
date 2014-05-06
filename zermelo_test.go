package ZermeloGo

import (
	"testing"
	"math/rand"
	"sort"
)

const TEST_SIZE = 100000

type uint32Sortable []uint32
type uint64Sortable []uint64

func (r uint32Sortable) Len() int           { return len(r) }
func (r uint32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r uint64Sortable) Len() int           { return len(r) }
func (r uint64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func TestSortUint32(t *testing.T) {
	var godata [TEST_SIZE]uint32
	g := godata[:]
	genTestDataUint32(g)
	var rdata [TEST_SIZE]uint32
	r := rdata[:]
	copy(r, g)
	sort.Sort(uint32Sortable(g))
	SortUint32(r)
	for i, val := range g {
		if r[i] != val {
			t.FailNow()
		}
	}
}

func TestSortUint64(t *testing.T) {
	var godata [TEST_SIZE]uint64
	g := godata[:]
	genTestDataUint64(g)
	var rdata [TEST_SIZE]uint64
	r := rdata[:]
	copy(r, g)
	sort.Sort(uint64Sortable(g))
	SortUint64(r)
	for i, val := range g {
		if r[i] != val {
			t.FailNow()
		}
	}
}

func genTestDataUint32(data []uint32) {
	for i, _ := range data {
		data[i] = rand.Uint32()
	}
}

func genTestDataUint64(data []uint64) {
	for i, _ := range data {
		data[i] = uint64(rand.Int63())
	}
}
