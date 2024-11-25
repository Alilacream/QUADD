package piscine

import (
	"github.com/01-edu/z01"
)

func QuadD(x, y int) {
	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			if j == 1 {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			} else if j > 0 && j < y {
				if i == 1 {
					z01.PrintRune('B')
				} else if i == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			} else if j == y {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
		}
		z01.PrintRune('\n')
	}
}