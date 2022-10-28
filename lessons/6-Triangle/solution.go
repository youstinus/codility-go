package main

import "sort"

// 100/100/100
func Solution(A []int) int {
	sort.SliceStable(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	for i := 0; i < len(A)-2; i++ {
		if A[i]+A[i+1] > A[i+2] && A[i+1]+A[i+2] > A[i] && A[i+2]+A[i] > A[i+1] {
			return 1
		}
	}

	return 0
}
