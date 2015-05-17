zermelo [![Build Status](https://travis-ci.org/shawnsmithdev/zermelo.svg)](https://travis-ci.org/shawnsmithdev/zermelo)  [![GoDoc](https://godoc.org/github.com/shawnsmithdev/zermelo?status.png)](https://godoc.org/github.com/shawnsmithdev/zermelo) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/shawnsmithdev/zermelo/master/LICENSE)
=========

A radix sorting library for Go.  Trade memory for speed!

```go
import "github.com/shawnsmithdev/zermelo"

func foo(large []uint64)
    zermelo.Sort(large)
}
```

About
------------

Zermelo is a sorting library featuring implementations of [radix sort](https://en.wikipedia.org/wiki/Radix_sort "Radix Sort"). I am especially influenced here by [these](http://codercorner.com/RadixSortRevisited.htm "Radix Sort Revisited") [two](http://stereopsis.com/radix.html "Radix Tricks") articles that describe various optimizations and how to work around the typical limitations of radix sort.

The code in general sacrifices DRY'ness for performance and an easy to use but flexible API. The general `Sort()` function works with any supported slice type, while more advanced options are available in subpackages for specific slice types.

Because this is a radix sort, it has a relatively large O(1) overhead costs in compute time, and will consume O(n) extra memory for the duration of the sort call. You will generally only want to use zermelo if your application is not memory constrained, and you will usually be sorting slices of supported types with at least 256 elements (128 for 32-bit types). The larger the slices you are sorting, the more benefit you will gain by using zermelo instead of the standard library's in-place comparison sort. The algorithm is also slightly adaptive; at certain times in the algorithm if the whole slice is found to be already sorted, no more work will be done. This is not as good as truely adaptive algorithms but will speed things up in certain cases.

### Etymology ###
Zermelo is named after [Ernst Zermelo](http://en.wikipedia.org/wiki/Ernst_Zermelo), who developed the proof for the well-ordered theorem, The theorem shows that sorting is possible, from a strict axiomatic mathmatical framework.

Zermelo Subpackages
-------------------
Using `zermelo.Sort()` allocates buffer space equal in size to the slice you are sorting, which must eventually be garbage collected. It also incurs a (very small) constant overhead for runtime reflection. While premature optimization should be avoided, this behavior may be a performance concern in demanding applications. Zermelo provides individual subpackages for each of the supported types, and new packages will be created as new types become supported. Subpackages have a `SortBYOB()` method where you can Bring Your Own Buffer (BYOB). Providing a buffer that is smaller than the slice you are sorting will cause a runtime panic. The subpackages `zuint` and `zint` must still use reflection as the bit width (32/64) is only available at runtime.

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
------

A Sorter will reuse buffers created during `Sort()` calls. This is not thread safe. Buffers are grown as needed at a 25% exponential growth rate.  This means if you sort a slice of size `n`, subsequent calls with slices up to `n * 1.25` in length will not cause another buffer allocation.

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

Benchmarks are not a promise of anything. You'll always want to profile for your use case.

You can run these on your own hardware

```Shell
go test -v -bench . -benchmem
```

Run with go 1.4.1 on a 2013 Macbook Air w/ i7-4650U and 8GB ram. For ns/op, lower is better.

[]uint64
--------

| slice len | golang ns/op | zermelo ns/op |Improvement|zermelo memory|
|-----------|--------------|---------------|-----------|--------------|
|64         |3783          |3617           | 4.39%     |  32  B       |
|256        |25839         |20707          |19.86%     |   2 KB       |
|65536      |14931449      |2593829        |82.63%     | 512 KB       |
|1048576    |298591046     |53842130       |81.97%     |   8 MB       |

[]float64
---------

| slice len | golang ns/op | zermelo ns/op |Improvement|zermelo memory|
|-----------|--------------|---------------|-----------|--------------|
|64         |6555          |6563           |-0.12%     |  32  B       |
|256        |41307         |24639          |49.87%     |   4 KB       |
|65536      |22999127      |3152232        |86.29%     |   1 MB       |
|1048576    |464524162     |58010014       |87.51%     |  16 MB       |

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
