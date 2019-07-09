// Filter using predicates.
package filter

import "reflect"

// Ints returns filtered slice of ints based on predicate.
func Ints(slice []int, fn func(v int) bool) []int {
	b := slice[:0]
	for _, v := range slice {
		if fn(v) {
			b = append(b, v)
		}
	}
	return b
}

// Float64s returns filtered slice of floats based on predicate.
func Float64s(slice []float64, fn func(v float64) bool) []float64 {
	b := slice[:0]
	for _, v := range slice {
		if fn(v) {
			b = append(b, v)
		}
	}
	return b
}

// Strings returns filtered slice of strings based on predicate.
func Strings(slice []string, fn func(v string) bool) []string {
	b := slice[:0]
	for _, v := range slice {
		if fn(v) {
			b = append(b, v)
		}
	}
	return b
}

// Slice filters in-place based on predicate.
func Slice(slice interface{}, fn func(i int) bool) {
	v := reflect.ValueOf(slice)
	rv := reflect.Indirect(v)

	if v.Kind() != reflect.Ptr || rv.Kind() != reflect.Slice {
		panic("filter: must be a pointer to slice")
	}

	length := rv.Len()
	for i := length - 1; i >= 0; i-- {
		if fn(i) {
			rv.Set(reflect.AppendSlice(rv.Slice(0, i), rv.Slice(i+1, rv.Len())))
		}
	}
}
