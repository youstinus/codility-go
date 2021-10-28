package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.DeepEqual(Solution([]int{0, 2, 3, 4, 1, 1, 1, 3, 1}), []int{1, 2, 3, 4, 1, 1, 1, 3, 2}))
	fmt.Println(reflect.DeepEqual(Solution([]int{0, 2, 3, 4, 1, 1, 1, 3, 1}), []int{0, 2, 4, 4, 1, 1, 2, 3, 1}), "alternative")
	fmt.Println(reflect.DeepEqual(Solution([]int{1, 1, 1, 2, 2, 1, 2, 2, 1}), []int{1, 1, 3, 2, 2, 1, 2, 2, 1}))
}
