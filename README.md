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

Lower is better

| Size  | # of keys | sort.Sort() ns/op | zermelo ns/op |
|-------|-----------|-------------------|---------------|
| Tiny  |512        |3408               |8669           |
| Small |2048       |22940              |22241          |
| nil   |524288     |13497885           |2249122        |
| Huge  |8388608    |269986939          |41310842       |

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
