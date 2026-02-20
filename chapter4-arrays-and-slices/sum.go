package main

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(arrNumbers ...[]int) []int {
	sums := []int{}
	for _, numbers := range arrNumbers {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(arrNumbers ...[]int) []int {
	sums := []int{}
	for _, numbers := range arrNumbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}
