zermelo/floats
==============
This subpackage handles sorting float slices. 

Example
-------

```go
package main

import (
	"github.com/shawnsmithdev/zermelo/floats"
	"something"
)

func main() {
	var x []float64
	x = something.GetFloatData()
	floats.SortFloats(x)
}
```

Sorter Example
--------------
todo