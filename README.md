# mu
### Go variadic math utilities for numeric types
Package mu provides variadic sum, min, max, abs functions for Go numeric
types. 
The name scheme is FuncXN where X is one of I, U, F and N is one of
8, 16, 32, 64, e,g.  `MinF32(nums ...float32) float32` returns the minimum of
its arguments which must all be of type `float32`.

There are a couple of exceptions to the naming rule:
- The funcs for the generic int type have no suffix, e.g. `Sum(nums ...int) int`.
- The `Abs` funcs are monadic e.g. `AbsI16(num) int16`.  Variadic versions are named VAbsXN, e.g. `VAbsI16(nums ...int16) []int16`

This package also provides a float64 comparison utility, `CloseEnough`.
