package piscine

func BasicAtoi(s string) int {
	var res int
	res = 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			res = res*10 + int(char-48)
		}
	}
	return res

}
