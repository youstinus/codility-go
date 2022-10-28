package main

// 60/100/20
func Solution(A []int) int {
	length := len(A)
	one := make([]float64, length)
	two := make([]float64, length)
	sumOne := 0.0
	sumTwo := 0.0

	for i := 0; i < length; i++ {
		sumOne += float64(A[i])
		sumTwo += float64(A[length-1-i])
		one[i] = sumOne
		two[length-1-i] = sumTwo
	}

	min := 2000000000.0
	index := -1

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			avg := ((one[j] + two[i]) - sumOne) / float64(j-i+1)
			if avg < min {
				min = avg
				index = i
			}
		}
	}

	return index
}

// 60/100/20
func Solution1(A []int) int {
	length := len(A)
	matrix := make([][]int, length-1)

	for i := 0; i < length-1; i++ {
		matrix[i] = make([]int, length-1-i)
		sum := A[i]
		for j := i + 1; j < length; j++ {
			sum += A[j]
			matrix[i][j-1-i] = sum
		}
	}

	min := 2000000000.0
	index := -1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			div := float64(matrix[i][j]) / float64(j+2)
			if div < min {
				min = div
				index = i
			}
		}
	}

	return index
}
