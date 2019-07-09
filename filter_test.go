package filter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	a := []int{6, 0, 1, 2, 3, 4, 5, 6}
	want := []int{6, 0, 2, 4, 6}

	Slice(&a, func(i int) bool {
		return a[i]%2 == 1
	})
	if !reflect.DeepEqual(a, want) {
		t.Fatalf("got %v want %v", a, want)
	}

	a = []int{6, 0, 1, 2, 3, 4, 5, 6}
	want = []int{1, 3, 5}

	Slice(&a, func(i int) bool {
		return a[i]%2 == 0
	})
	if !reflect.DeepEqual(a, want) {
		t.Fatalf("got %v want %v", a, want)
	}

}
