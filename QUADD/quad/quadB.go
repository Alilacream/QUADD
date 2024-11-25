package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {


			if j == 1 {
				if i == 1 {
					z01.PrintRune('/')
				} else if i == x {
					z01.PrintRune('\\')
				} else {
					z01.PrintRune('*')
				}
			} else if j > 1 && j < y {
				if i == 1 {
					z01.PrintRune('*')
				} else if i == x {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			} else if j == y {
				if i == 1 {
					z01.PrintRune('\\')
				} else if i == x {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			}
		}
		z01.PrintRune('\n')
	}
}