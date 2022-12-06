package functional_test

import (
	f "github.com/rangzen/go-boataootz/functional"
	"github.com/rangzen/go-boataootz/test"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("addition of all elements", func(t *testing.T) {
		data := []int{1, 2, 3}
		acc := func(a, b int) int {
			return a + b
		}
		init := 0
		want := 6

		got := f.Reduce(data, acc, init)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		test.AssertEqual(t, f.Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		test.AssertEqual(t, f.Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}
