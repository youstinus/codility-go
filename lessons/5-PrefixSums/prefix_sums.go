package main

import "math"

// Mushrooms is a demo from pdf reading material
func Mushrooms(A []int, k int, m int) int {
	n := len(A)
	result := 0
	pref := prefixSums(A)
	for p := 0; p < int(math.Min(float64(m), float64(k)))+1; p++ {
		leftPos := k - p
		rightPos := int(math.Min(float64(n-1), math.Max(float64(k), float64(k+m-2*p))))
		result = int(math.Max(float64(result), countTotal(pref, leftPos, rightPos)))
	}

	for p := 0; p < int(math.Min(float64(m+1), float64(n-k))); p++ {
		rightPos := k + p
		leftPos := int(math.Max(0, math.Min(float64(k), float64(k-(m-2*p)))))
		result = int(math.Max(float64(result), countTotal(pref, leftPos, rightPos)))
	}
	return result
}

func prefixSums(A []int) []int {
	n := len(A)
	P := make([]int, n+1)
	for k := 1; k < n+1; k++ {
		P[k] = P[k-1] + A[k-1]
	}
	return P
}

func countTotal(P []int, x int, y int) float64 {
	return float64(P[y+1] - P[x])
}
