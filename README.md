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

Run on a 2013 Macbook Air w/ i7-4650U and 8GB ram.  Lower is better.

| Size  | # of keys | sort.Sort() ns/op | zermelo ns/op |
|-------|-----------|-------------------|---------------|
| Tiny  |64         |3606               |3765           |
| Small |256        |24844              |13976          |
| nil   |65536      |15509938           |2327765        |
| Huge  |1048576    |309781344          |50645600       |

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
