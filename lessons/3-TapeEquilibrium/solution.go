package main

// 100/100/100
func Solution(A []int) int {
	sum := 0
	for _, a := range A {
		sum += a
	}
	min := 100000000000
	sum1 := 0
	for i := 0; i < len(A)-1; i++ {
		a := A[i]
		sum1 += a
		absolute := abs(sum - sum1 - sum1)
		if absolute < min {
			min = absolute
		}
	}

	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
