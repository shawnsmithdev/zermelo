package zermelo

import "testing"

func TestUnsafeSliceConvert(t *testing.T) {
	foo := []float64{0, 3.14, 88.888, -1}
	bar := unsafeSliceConvert[float64, []float64, uint64](foo)
	if len(foo) != len(bar) {
		t.Fatal("foo-bar len")
	}
	if cap(foo) != cap(bar) {
		t.Fatal("foo-bar cap")
	}
}
