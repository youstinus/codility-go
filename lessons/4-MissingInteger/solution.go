package main

import "sort"

// 100/100/100
func Solution(A []int) int {
	sort.SliceStable(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	if A[len(A)-1] < 1 {
		return 1
	}

	if A[0] > 1 {
		return 1
	}

	prev := A[0]
	if prev < 0 {
		prev = 0
	}
	for _, a := range A {
		if a < 1 {
			continue
		}
		if a == prev {
			continue
		} else if prev+1 == a {
			prev = a
		} else {
			return prev + 1
		}
	}

	return prev + 1
}
