package piscine

func UltimateDivMod(a *int, b *int) {
	s1 := *a
	s2 := *b
	*a = s1 / s2
	*b = s1 % s2

}
