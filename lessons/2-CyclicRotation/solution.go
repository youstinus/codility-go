package main

// 100/100/100
func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	if K == 0 {
		return A
	}

	length := len(A)
	K = K % length
	B := make([]int, length)

	for i := 0; i < length; i++ {
		B[(K+i)%length] = A[i]
	}

	return B
}
