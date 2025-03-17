package array_sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array sum of size 5", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		got := Sum(arr)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, arr)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of any number of arrays, should return array of their respective sum", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2})
		want := []int{6, 15, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSum(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})

}
