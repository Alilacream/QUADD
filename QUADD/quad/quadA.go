package piscine

import "github.com/01-edu/z01"

func Quada(x, y int) {
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			if ((i == 0) && ((j == 0) || (j == (y - 1)))) || ((i == x-1) && ((j == 0) || (j == y-1))) {
				z01.PrintRune('o')
			} else if ((j > 0) && (j < (y - 1))) && ((i == 0) || (i == (x - 1))) {
				z01.PrintRune('|')
			} else if (i != 0) && (i != (x - 1)) && (j != 0) && (j != (y - 1)) {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
	}
}
