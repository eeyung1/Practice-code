package main


import (
	"fmt"
)

func main(){
	for i := 1; i <= 9; i++ {
		for j := 9; j >= 0 ; j-- {
			fmt.Printf("%d%d ", i, j)
		}
	}
}

/*

func main() {
	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				if i == 2 && j == 1 && k == 0 {
					fmt.Printf("%d%d%d\n", i, j, k)
				} else {
					fmt.Printf("%d%d%d, ", i, j, k)
				}
			}
		}
	}
}
*/


/*

func PrintRevComb() {
	for i := '9'; i >= '2'; i-- {
		for j := i - 1; j >= '1'; j-- {
			for k := j - 1; k >= '0'; k-- {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				// Print comma and space for all except the last one (210)
				if !(i == '2' && j == '1' && k == '0') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
	*/
