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


| Size  | Bytes | Kilobytes | Megabytes| sort.Sort() iter/1s | sort.Sort() time | zermelo iter/1s | zermelo time   |
|-------|-------|-----------|----------|---------------------|------------------|-----------------|----------------|
| Tiny  |512    |   0.5     |          |              500000 |      3408 ns/op  |          200000 |     8669 ns/op |
| Small |2048   |   2       |          |              100000 |     22940 ns/op  |          100000 |    22241 ns/op |
| nil   |524288 |   512     |   0.5    |                 100 |  13497885 ns/op  |            1000 |  2249122 ns/op |
| Huge  |8388608|   8192    |   8      |                   5 | 269986939 ns/op  |              50 | 41310842 ns/op |

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

