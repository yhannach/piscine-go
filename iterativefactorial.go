package piscine

/*func IterativeFactorial(nb int) int {
	res := 1

	for i := 1; i <= nb; i++ {
		res = res * i
	}

	return res
} */

func IterativeFactorial(nb int) int {

	var itera int
	itera = 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 24 {
		for i := 1; i <= nb; i++ {
			itera = itera * i
		}
	} else {
		return 0
	}

	return itera
}
