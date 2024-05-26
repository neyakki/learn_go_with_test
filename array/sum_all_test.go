package array

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestSumAllTail(t *testing.T) {
	t.Run("not empty slice", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{0, 9})
		expected := []int{0, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
}
