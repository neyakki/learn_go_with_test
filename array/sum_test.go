package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		var numbers []int = []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		var expected int = 15

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})
}
