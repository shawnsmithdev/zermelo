zermelo [![Build Status](https://travis-ci.org/shawnsmithdev/zermelo.svg)](https://travis-ci.org/shawnsmithdev/zermelo) [![GoDoc](https://godoc.org/github.com/shawnsmithdev/zermelo?status.png)](https://godoc.org/github.com/shawnsmithdev/zermelo)
=========

A performance sorting library for Golang.

```go
import "github.com/shawnsmithdev/zermelo"

func foo(large []uint64)
    zermelo.SortUint64(large)
}
```

uint64 Benchmarks
-----------------

Size chart

|Name | Bytes | Kilobytes | Megabytes|
|-----|-------|-----------|----------|
|Tiny |512    |   0.5     |          |
|Small|2048   |   2       |          |
|     |524288 |   512     |   0.5    |
|Huge |8388608|   8192    |   8      |

Results

| Benchmark(lib)Sort(type)(size)  | iter/s  |      speed      |
|---------------------------------|---------|-----------------|
| BenchmarkZermeloSortUint64Tiny  | 200000  |      8669 ns/op |
| BenchmarkGoSortUint64Tiny       | 500000  |      3408 ns/op |
| BenchmarkZermeloSortUint64Small | 100000  |     22241 ns/op |
| BenchmarkGoSortUint64Small      | 100000  |     22940 ns/op |
| BenchmarkZermeloSortUint64      |   1000  |   2249122 ns/op |
| BenchmarkGoSortUint64           |    100  |  13497885 ns/op |
| BenchmarkZermeloSortUint64Big   |     50  |  41310842 ns/op |
| BenchmarkGoSortUint64Big        |      5  | 269986939 ns/op |

Working
-------

* []uint32
* []uint64

TODO
----

* ~~Benchmarks~~ done
* Split into files by type
* Signed ints
* Floats
* Sort() call that uses O(1) reflection with sort.Sort() base case
* Move type specific code to subpackages

