zermelo/floats
==============
This subpackage handles sorting float slices. 

Example
-------

```go
package main

import (
	"github.com/shawnsmithdev/zermelo/v2/floats"
	"something"
)

func main() {
	var x []float64
	x = something.GetFloatData()
	floats.SortFloats(x)
}
```

Sorter
======

The `Sorter` returned by `NewFloatSorter()` will reuse buffers created during `Sort()` calls. This is not thread safe,
and behaves in the same manner as `zermelo.NewSorter()`, but for float types.

Sorter Example
--------------
```go
package main

import (
	"github.com/shawnsmithdev/zermelo/v2/floats"
	"something"
)

func main() {
	var x [][]float64
	x = something.GetFloatDatas()
	sorter := floats.NewFloatSorter[float64]()
	for _, y := range x {
		sorter.Sort(y)
    }
}
```
