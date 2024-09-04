package main

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(numSlices ...[]int) []int {
	var sums []int

	for _, slice := range numSlices {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(numSlices ...[]int) []int {
	var sums []int

	for _, slice := range numSlices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}
	return sums
}