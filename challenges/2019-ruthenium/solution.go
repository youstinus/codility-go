package main

// Solution exported
// award and performance here
func Solution(A []int, K int) int {

	max := 0
	k := K
	for i := 0; i < len(A)-1; i++ {
		sum := 1
		for j := i + 1; j < len(A); j++ {
			if A[i] != A[j] {
				if k <= 0 {
					break
				}
				k--
			}
			sum++
		}
		k = K
		if sum > max {
			max = sum
		}
	}

	return max
}
