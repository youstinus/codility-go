package main

// 75/75/75
func Solution(N int, A []int) []int {
	B := make([]int, N)
	maxVal := 0
	for _, a := range A {
		if 1 <= a && a <= N {
			B[a-1]++
			if maxVal < B[a-1] {
				maxVal = B[a-1]
			}
		} else if a == N+1 {
			setMax(B, maxVal)
		}
	}
	return B
}

func setMax(A []int, maxVal int) {
	for i := range A {
		A[i] = maxVal
	}
}
