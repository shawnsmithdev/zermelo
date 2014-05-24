zermelo [![Build Status](https://travis-ci.org/shawnsmithdev/zermelo.svg)](https://travis-ci.org/shawnsmithdev/zermelo) [![GoDoc](https://godoc.org/github.com/shawnsmithdev/zermelo?status.png)](https://godoc.org/github.com/shawnsmithdev/zermelo)
=========

A performance sorting library for Golang.

```go
import "github.com/shawnsmithdev/zermelo"

func foo(large []uint64)
    zermelo.SortUint64(large)
}
```

Design Goals
------------

Overall these sort implementations will utilize a [radix sort](https://en.wikipedia.org/wiki/Radix_sort "Radix Sort").
I am especially influenced here by the [these](http://codercorner.com/RadixSortRevisited.htm "Radix Sort Revisited")
[two](http://stereopsis.com/radix.html "Radix Tricks") articles that describe various optimizations and how
to work around the typical limitations of radix sort.

The code will in general sacrifice DRY'ness for performance and a clean external API.  It is intended for there to be a simple, general Sort() function that applies to all types and sizes, and more advanced options to eek out a bit more performance when you know a lot about the data you are sorting.

Because this is a radix sort, it has a relatively large O(n) overhead costs in both compute time and memory.
The general Sort() function will also have some O(1) reflection overhead.  You will generally only want to
use zermelo if you know that your application is not memory constrained, and you will usually be sorting slices
of supported types with at least 256 elements.  The larger the slices you are sorting, the more benefit you will
gain by using zermelo instead of the traditionally approach of aliasing the slice type to a Sortable type
and using sort.Sort().


uint64 Benchmarks
-----------------

Run on a 2013 Macbook Air w/ i7-4650U and 8GB ram.  Lower is better.

| Size  | # of keys | sort.Sort() ns/op | zermelo ns/op |Improvement|
|-------|-----------|-------------------|---------------|-----------|
| Tiny  |64         |3606               |3765           |-4.41%     |
| Small |256        |24844              |13976          |43.74%     |
| nil   |65536      |15509938           |2327765        |84.99%     |
| Huge  |1048576    |309781344          |50645600       |83.65%     |

Working
-------

* []uint32
* []uint64
* []int32

TODO
----

* ~~Benchmarks~~ done
* ~~Split into files by type~~ done
* Signed ~~int32~~, int64, int
* Floats
* Sort() call that uses O(1) reflection with sort.Sort() base case
* Move type specific code to subpackages
