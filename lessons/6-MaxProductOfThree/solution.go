package main

import "sort"

// 100/100/100
func Solution(A []int) int {
	sort.SliceStable(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	topMax := A[0] * A[1] * A[2]
	lowMax := A[0] * A[len(A)-1] * A[len(A)-2]
	if topMax > lowMax {
		return topMax
	}

	return lowMax
}

// 44/100/0
func Solution1(A []int) int {
	max := -2000000000
	for i := 0; i < len(A)-2; i++ {
		for j := i + 1; j < len(A)-1; j++ {
			for k := j + 1; k < len(A); k++ {
				triplet := A[i] * A[j] * A[k]
				if triplet > max {
					max = triplet
				}
			}
		}
	}

	return max
}
