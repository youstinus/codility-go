package main

import (
	"fmt"
)

func main() {
	A := []int{2, 3, 7, 5, 1, 3, 9}
	result := Mushrooms(A, 4, 6)
	fmt.Println(result)

	B := []int{0, 1, 0, 1, 1}
	result1 := PassingCarsSolution(B)
	fmt.Println(result1)

}
