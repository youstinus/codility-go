package main

import "sort"

// Solution gets solution
func Solution(A []int) int {
	sort.Ints(A)
	//var B = make([]int, 0)
	i := 1
	for _, item := range A {
		if item > 0 {
			if item == i {
				i++
			} else if item > 0 && item > i {
				return i
			}
		}
	}

	return 1
}
