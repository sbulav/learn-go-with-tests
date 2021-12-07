package arrays

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number

	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToTail ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToTail {
		if len(numbers) > 0 {
			sums = append(sums, Sum(numbers[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
