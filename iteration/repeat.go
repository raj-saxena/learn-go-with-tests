package iteration

func Repeat(char string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += char
	}
	return result
}
