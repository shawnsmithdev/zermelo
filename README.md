zermelo [![Build Status](https://travis-ci.org/shawnsmithdev/zermelo.svg)](https://travis-ci.org/shawnsmithdev/zermelo) [![GoDoc](https://godoc.org/github.com/shawnsmithdev/zermelo?status.png)](https://godoc.org/github.com/shawnsmithdev/zermelo)
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

The code will in general sacrifice DRY'ness for performance and a clean external API.  There is a general Sort() function that applies to all types and sizes, and more advanced options to avoid reflection when you know the type of the data you are sorting.

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

uint64 Benchmarks
-----------------

Run on a 2013 Macbook Air w/ i7-4650U and 8GB ram. For ns/op, lower is better.

| Size  | # of keys | sort.Sort() ns/op | zermelo ns/op |Improvement|
|-------|-----------|-------------------|---------------|-----------|
| Tiny  |64         |4361               |4631           |-6.19%     |
| Small |256        |28390              |20938          |26.25%     |
| nil   |65536      |17187801           |2856613        |83.38%     |
| Huge  |1048576    |343859473          |59408405       |82.72%     |

Working
-------

* []uint32
* []uint64
* []int32
* []int64
* int[]
* uint[]
* Move type specific code to subpackages (ex. github.com/shawnsmithdev/zermelo/zint64)

TODO
----

* float32[]
* float64[]
