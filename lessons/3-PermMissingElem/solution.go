package main

// 100/100/100
func Solution(A []int) int {
	length := len(A)
	sum := 0
	calc := (length + 1) * (2 + length) / 2

	for _, a := range A {
		sum += a
	}

	return calc - sum
}
