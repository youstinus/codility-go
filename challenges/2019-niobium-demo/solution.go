package main

import (
	"sort"
)

// Solution exported
func Solution(A []int) int {
	A = unique(A)
	sort.Ints(A)

	for i := 1; i <= len(A); i++ {
		if i == A[i-1] {
			continue
		}
		return i
	}

	if len(A) == 0 {
		return 1
	}

	return A[len(A)-1] + 1
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value && entry > 0 {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
