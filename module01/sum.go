package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	cnt := 0

	for _, num := range numbers {
		cnt += num
	}

	return cnt
}
