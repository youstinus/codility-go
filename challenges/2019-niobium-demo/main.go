package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3} //{1, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5}
	var result = Solution(A)
	fmt.Println(result)
}
