package int

func Pow(base int, exponent uint) int {
	answer := 1
	for i := 0; i < int(exponent); i++ {
		answer *= base
	}
	return answer
}
