//go:generate ./gen_mu.py
//go:generate go fmt

// Package mu provides variadic sum, min, max, abs functions for Go numeric
// types. The name scheme is FuncXN where X is one of I, U, F and N is one of
// 8, 16, 32, 64, e,g.  MinF32(nums ...float32) float32 returns the minimum of
// its arguments which must all be of type float32.
// There are a couple of exceptions to the rule:
//
// The funcs for the generic int type have no suffix, e.g. Sum(nums ...int) int.
//
// The Abs funcs are monadic e.g. AbsI16(num) int16.  Variadic versions are
// named VAbsXN, e.g. VAbsI16(nums ...int16) []int16
//
// This package also provides a float64 comparison utility, CloseEnough.
package mu

import "math"

// CloseEnough compares two float64. Tolerance is the minimum acceptable
// ratio of the difference between a and b relative to their average absolute value.
func CloseEnough(a, b, tolerance float64) bool {
	if a == 0 && b == 0 {
		return true // avoids div 0 error
	}
	// at least one value is non-zero
	denom := 2 * math.Abs(a-b)
	numer := math.Abs(a) + math.Abs(b)
	return math.Abs(tolerance) > denom/numer
}
