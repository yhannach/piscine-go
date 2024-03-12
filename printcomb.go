package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for y := '0'; y <= '9'; y++ {
				if i < j && j < y {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(y)
					if !(i == '7' && j == '8' && y == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}
