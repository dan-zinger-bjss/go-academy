package iteration

func Repeat(input string, iterations int) (output string) {
	if iterations == -1 {
		iterations = 5
	}
	for i := 0; i < iterations; i ++ {
		output += input
	}
	return output
}