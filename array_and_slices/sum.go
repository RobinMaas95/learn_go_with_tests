package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, arr := range numbersToSum {
		sums = append(sums, Sum(arr))
	}
	return sums
}
