package zermelo

// sort.Sort()'able types
type uint32Sortable []uint32
type uint64Sortable []uint64
type int32Sortable []int32
type int64Sortable []int64

// implements sort interface
func (r uint32Sortable) Len() int           { return len(r) }
func (r uint32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r uint64Sortable) Len() int           { return len(r) }
func (r uint64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r int32Sortable) Len() int           { return len(r) }
func (r int32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r int64Sortable) Len() int           { return len(r) }
func (r int64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
