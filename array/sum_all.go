package array

func SumAll(numbers ...[]int) []int {
	var sums []int

	for index := range numbers {
		sums = append(sums, Sum(numbers[index]))
	}

	return sums
}

func SumAllTail(numbers ...[]int) []int {
	sums := make([]int, 0)

	for index := range numbers {
		if len(numbers[index]) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[index][1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
