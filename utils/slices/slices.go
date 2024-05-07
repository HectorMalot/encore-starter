package slices

import (
	"cmp"
	goslices "slices"
)

// Map applies a function to each element of a slice and returns a new slice
// containing the results.
func Map[T1, T2 any](in []T1, fn func(T1) T2) []T2 {
	out := make([]T2, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

// Filter applies a function to each element of a slice and returns a new slice
// containing all elements for which the function returns true.
func Filter[T any](in []T, fn func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

// Unique returns a new slice containing only the unique elements of the input
// slice.
func Unique[T cmp.Ordered](in []T) []T {
	goslices.Sort(in)
	out := goslices.Compact(in)
	return out
}
