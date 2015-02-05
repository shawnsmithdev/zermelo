zermelo [![Build Status](https://travis-ci.org/shawnsmithdev/zermelo.svg)](https://travis-ci.org/shawnsmithdev/zermelo)  [![GoDoc](https://godoc.org/github.com/shawnsmithdev/zermelo?status.png)](https://godoc.org/github.com/shawnsmithdev/zermelo) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/shawnsmithdev/zermelo/master/LICENSE)
=========

A performance sorting library for Golang.

```go
import "github.com/shawnsmithdev/zermelo"

func foo(large []uint64)
    zermelo.Sort(large)
}
```

Design Goals
------------

Overall these sort implementations will utilize a [radix sort](https://en.wikipedia.org/wiki/Radix_sort "Radix Sort").
I am especially influenced here by [these](http://codercorner.com/RadixSortRevisited.htm "Radix Sort Revisited")
[two](http://stereopsis.com/radix.html "Radix Tricks") articles that describe various optimizations and how
to work around the typical limitations of radix sort.

The code will in general sacrifice DRY'ness for performance and a clean external API.  There is a general Sort() function that applies to all types and sizes, while more advanced options are available in subpackages to avoid reflection when you know the type of the data you are sorting.

Because this is a radix sort, it has a relatively large O(1) overhead costs in compute time, moreso with reflection, and will consume O(n) extra memory for the duration of the sort call. You will generally only want to use zermelo if you know that your application is not memory constrained, and you will usually be sorting slices of supported types with at least 256 elements. The larger the slices you are sorting, the more benefit you will gain by using zermelo instead of the traditionally approach of aliasing the slice type to a Sortable type and using sort.Sort().

The sort is not adaptive in the traditional sense, but I plan to implement a check to short circuit a lot of the work if it is detected that the slice is already sorted.  Stability is not relevant as zermelo only supports slices of numeric types (except the general Sort() method with sort.Sortable types that are not numeric slices, as those will be sorted by the standard library's comparison sort, which is stable).

Zermelo Subpackages
-------------------
Using zermelo.Sort() incurs a small constant overhead for runtime reflection.  It also allocates buffer space, which must eventually be garbage collected. While premature optimization should be avoided, this behavior may be a performance concern in demanding applications. Zermelo provides individual subpackages for each of the supported types, and new packages will be created as new types become supported.

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

Benchmarks
==========

Benchmarks are not a promise of anything. You'll always want to profile for your use case.

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
|256        |41307         |24639          |15.96%     |   4 KB       |
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
