zermelo [![Build Status](https://travis-ci.org/shawnsmithdev/zermelo.svg)](https://travis-ci.org/shawnsmithdev/zermelo)  [![GoDoc](https://godoc.org/github.com/shawnsmithdev/zermelo?status.png)](https://godoc.org/github.com/shawnsmithdev/zermelo) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/shawnsmithdev/zermelo/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/shawnsmithdev/zermelo)](https://goreportcard.com/report/github.com/shawnsmithdev/zermelo)  [![Sourcegraph](https://sourcegraph.com/github.com/shawnsmithdev/zermelo/-/badge.svg)](https://sourcegraph.com/github.com/shawnsmithdev/zermelo?badge)
=========

A radix sorting library for Go.  Trade memory for speed!

```go
import "github.com/shawnsmithdev/zermelo"

func foo(large []uint64)
    zermelo.Sort(large)
}
```

About
=====

Zermelo is a sorting library featuring implementations of [radix sort](https://en.wikipedia.org/wiki/Radix_sort "Radix Sort"). I am especially influenced here by [these](http://codercorner.com/RadixSortRevisited.htm "Radix Sort Revisited") [two](http://stereopsis.com/radix.html "Radix Tricks") articles that describe various optimizations and how to work around the typical limitations of radix sort.

You will generally only want to use zermelo if you won't mind the extra memory used for buffers and your application frequently sorts slices of supported types with at least 256 elements (128 for 32-bit types). The larger the slices you are sorting, the more benefit you will gain by using zermelo instead of the standard library's in-place comparison sort.

Etymology
---------
Zermelo is named after [Ernst Zermelo](http://en.wikipedia.org/wiki/Ernst_Zermelo), who developed the proof for the [well-ordering theorem](https://en.wikipedia.org/wiki/Well-ordering_theorem).

Supported Types
===============
* []float32
* []float64
* []int
* []int32
* []int64
* []uint
* []uint32
* []uint64

Subpackages
===========
Zermelo provides individual subpackages for each of the supported types. Subpackages have a `SortBYOB()` method where you can Bring Your Own Buffer (BYOB), for minimizing allocations. Providing a buffer that is smaller than the slice you are sorting will cause a runtime panic.

```go
import "github.com/shawnsmithdev/zermelo/zuint64"

func foo(bar SomeRemoteData)
    data := make([]uint64, REALLY_BIG)
    buffer := make([]uint64, REALLY_BIG)

    while bar.hasMore() {
        bar.Read(data)
        zuint64.SortBYOB(data, buffer)
        doSomething(data)
    }
}
```

Sorter
======

A Sorter will reuse buffers created during `Sort()` calls. This is not thread safe. Buffers are grown as needed at a 25% exponential growth rate.  This means if you sort a slice of size `n`, subsequent calls with slices up to `n * 1.25` in length will not cause another buffer allocation. This does not apply to the first allocation, which will make a buffer of the same size as the requested slice. This way, if the slices being sorted do not grow in size, there is no unused buffer space.

```go
import "github.com/shawnsmithdev/zermelo"

func foo(bar [][]uint64) {
    sorter := zermelo.New()
    for _, x := range(bar) {
        sorter.Sort(x)
    }
}

```

Benchmarks
==========

You can run the benchmark on your own hardware.

```Shell
go test -v -bench . -benchmem
```

The benchmark tests two types of slices:
* `[]uint64`
* `[]float64`

For each of the tested types, it runs the benchmark on a slice of that type with four sizes:
* `T` (tiny) 64
* `S` (small) 256
* `M` (medium) 1024
* `L` (large) 1048576

For each slice type and size, three sorters are benchmarked:
* GoSort - The standard library sort: `sort.Slice()` or `sort.Float64s`
* ZSort - `zermelo.Sort()`, does not reuse buffers
* ZSorter - `zermelo.New().Sort()`, does reuse buffers

One pass for each sorter is also made against a large, presorted `[]uint64`.
