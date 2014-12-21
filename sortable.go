package zermelo

// Declare sortable type aliases for sort.Sort()
type uintSortable []uint
type uint32Sortable []uint32
type uint64Sortable []uint64
// int[] provided by sort.IntSlice
type int32Sortable []int32
type int64Sortable []int64
type float32Sortable []float32
// float64[] provided by sort.Float64Slice

// implements sort interface
// uint[]
func (r uintSortable) Len() int           { return len(r) }
func (r uintSortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uintSortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
// uint32[]
func (r uint32Sortable) Len() int           { return len(r) }
func (r uint32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
// uint64[]
func (r uint64Sortable) Len() int           { return len(r) }
func (r uint64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
// int[] provided by sort.IntSlice
// int32[]
func (r int32Sortable) Len() int           { return len(r) }
func (r int32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
// int64[]
func (r int64Sortable) Len() int           { return len(r) }
func (r int64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
// float32[]
func (r float32Sortable) Len() int           { return len(r) }
func (r float32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r float32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
// float64[] provided by sort.Float64Slice

