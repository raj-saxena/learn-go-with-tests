package arrayAndSlices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	got := Sum(values)
	want := 28

	if got != want {
		t.Errorf("got %v, want %v, given %v", got, want, values)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Sum tails of slices", func(t *testing.T) {
		got := SumTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumTails([]int{}, []int{4, 5, 6})
		want := []int{0, 11}

		checkSums(t, got, want)
	})
}
