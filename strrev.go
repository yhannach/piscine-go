package piscine

func StrRev(s string) string {
	var str string
	for i := len(s) - 1; i >= 0; i-- {
		str = str + string(s[i])
	}
	return str
}
