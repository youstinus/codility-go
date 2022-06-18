package main

// 100/100/100
func Solution(A []int) int {
	sum := len(A)
	// sevens := make([]int, sum)
	in1, in2 := 0, 0

	for {
		if in1 >= len(A) || in2 >= len(A) {
			break
		}

		diff := A[in2] - A[in1]
		if diff > 6 {
			in1++
		} else if diff == 6 {

		} else if diff < 6 {
			in2++
		}
	}

	return sum
}
