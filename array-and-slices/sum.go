package arrayAndSlices

func Sum(values []int) (sum int) {
	for _, v := range values {
		sum += v
	}
	return
}

func SumAll(values ...[]int) []int {
	resultSize := len(values)
	result := make([]int, resultSize)

	for i, num := range values {
		result[i] = Sum(num)
	}
	return result
}

func SumTails(values ...[]int) []int {
	var result []int

	for _, num := range values {
		if len(num) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(num[1:]))
		}
	}
	return result
}
