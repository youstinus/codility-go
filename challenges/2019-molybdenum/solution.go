package main

import (
	"sort"
)

// Solution exported
// Silver award. 100% correct, but score 55, performance 0
func Solution(K int, M int, A []int) []int {
	B := preparation(A, M)
	return proceed(A, B, K)
}

func proceed(A []int, B []int, K int) []int {
	C := make([]int, 0)

	for i := 0; i <= len(A)-K; i++ {
		for k := i; k < i+K; k++ {
			B[A[k]]--
			B[A[k]+1]++
		}

		for i1, v := range B {
			if 2*v > len(A) && !contains(C, i1) {
				C = append(C, i1)
			}
		}

		for k := i; k < i+K; k++ {
			B[A[k]]++
			B[A[k]+1]--
		}

	}

	sort.Ints(C)

	return C
}

func contains(C []int, v int) bool {
	for _, v1 := range C {
		if v1 == v {
			return true
		}
	}
	return false
}

func preparation(A []int, M int) []int {
	B := make([]int, M+2) // max(A)
	for _, c := range A {
		B[c]++
	}
	return B
}
