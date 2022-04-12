zermelo [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/shawnsmithdev/zermelo) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/shawnsmithdev/zermelo/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/shawnsmithdev/zermelo)](https://goreportcard.com/report/github.com/shawnsmithdev/zermelo)
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

You will generally only want to use zermelo if you won't mind the extra memory used for buffers and your application frequently sorts slices of supported types with at least 256 elements (128 for 32-bit types, somewhat more for floats on ARM). The larger the slices you are sorting, the more benefit you will gain by using zermelo instead of the standard library's in-place comparison sort.

Etymology
---------
Zermelo is named after [Ernst Zermelo](http://en.wikipedia.org/wiki/Ernst_Zermelo), who developed the proof for the [well-ordering theorem](https://en.wikipedia.org/wiki/Well-ordering_theorem).

Supported Types
===============
constraints.Integer and constraints.Float

Sorter
======

A Sorter will reuse buffers created during `Sort()` calls. This is not thread safe. Buffers are grown as needed at a 25% exponential growth rate.  This means if you sort a slice of size `n`, subsequent calls with slices up to `n * 1.25` in length will not cause another buffer allocation. This does not apply to the first allocation, which will make a buffer of the same size as the requested slice. This way, if the slices being sorted do not grow in size, there is no unused buffer space.

```go
import "github.com/shawnsmithdev/zermelo"

func foo(bar [][]uint64) {
    sorter := zermelo.NewIntSorter[uint64]()
    for _, x := range(bar) {
        sorter.Sort(x)
    }
}

```