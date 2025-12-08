package main

import "fmt"

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}
	n := len(numbers)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	medianValue := numbers[2]

	return medianValue
}

func main(){
fmt.Println(Abort(2,3,3,4,4))

}
