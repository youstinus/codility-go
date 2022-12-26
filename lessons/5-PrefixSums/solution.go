package main

// Solution 100/100/100
func Solution(A []int) int {
	suf := suffixSum(A)
	sum := 0
	for i, v := range A {
		if v == 0 {
			sum += suf[i+1]
		}
	}

	if sum > 1000000000 {
		return -1
	}
	return sum
}

func suffixSum(A []int) []int {
	n := len(A)
	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {

		sum[i] = sum[i+1] + A[i]
	}
	return sum
}
