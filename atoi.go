package piscine

func Atoi(s string) int {
	sin := 1
	res := 0

	for i, v := range s {
		if v >= '0' && v <= '9' {
			res = res*10 + int(v-48)
		} else if i == 0 && v == '-' {
			sin = -1
		} else if i == 0 && v == '+' {
			sin = 1
		} else {
			return 0
		}

	}
	return res * sin
}
