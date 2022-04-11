package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, arr := range numbersToSum {
		var appender int
		if len(arr) < 1 {
			appender = 0
		} else {
			appender = Sum(arr[1:])
		}
		sums = append(sums, appender)
	}
	return sums
}
