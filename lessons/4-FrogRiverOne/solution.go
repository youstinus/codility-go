package main

// 100/100/100
func Solution(X int, A []int) int {
	leaves := make([]int, X)
	sum := 0
	for i, a := range A {
		old := leaves[a-1]
		if old == 0 {
			sum++
		}
		leaves[a-1] = 1

		if sum == X {
			return i
		}
	}

	return -1
}
