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
I am especially influenced here by [these](http://codercorner.com/RadixSortRevisited.htm "Radix Sort Revisited")
[two](http://stereopsis.com/radix.html "Radix Tricks") articles that describe various optimizations and how
to work around the typical limitations of radix sort.

The code will in general sacrifice DRY'ness for performance and a clean external API.  It is intended for there to be a simple, general Sort() function that applies to all types and sizes, and more advanced options to eek out a bit more performance when you know a lot about the data you are sorting.

Because this is a radix sort, it has a relatively large O(1) overhead costs in compute time, and will
consume O(n) extra memory for the duration of the sort call. The general Sort() function will also have
some O(1) reflection overhead.  You will generally only want to use zermelo if you know that your application
is not memory constrained, and you will usually be sorting slices of supported types with at least 256 elements.
The larger the slices you are sorting, the more benefit you will gain by using zermelo instead of the
traditionally approach of aliasing the slice type to a Sortable type and using sort.Sort().

The sort is not adaptive in the traditional sense, but I plan to implement a check to short circuit a lot of the work if it is detected that the slice is already sorted.  Stability is not relevant as zermelo only supports slices of numeric types (except the general Sort() method with sort.Sortable types that are not numeric slices, as those will be sorted by the standard library's comparison sort, which is stable).

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
* []int64

TODO
----

* ~~Benchmarks~~ done
* ~~Split into files by type~~ done
* Signed ~~int32~~, ~~int64~~, int
* Floats
* Sort() call that uses O(1) reflection with sort.Sort() base case
* Move type specific code to subpackages
