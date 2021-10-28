package main

import "fmt"

// Solution exported
// 100/100/100
func Solution(A []int) []int {

	sums := make([]int, 6)

	sums[0] = A[0] + A[1] + A[2]
	sums[1] = A[3] + A[4] + A[5]
	sums[2] = A[6] + A[7] + A[8]
	sums[3] = A[0] + A[3] + A[6]
	sums[4] = A[1] + A[4] + A[7]
	sums[5] = A[2] + A[5] + A[8]

	m := max(sums)
	inc := make([]int, 6)

	for i, s := range sums {
		inc[i] = m - s
	}

	if !isValid(inc) {
		fmt.Println("validation failed")
		return A
	}

	i := 0
	j := 3
	for {
		if i > 2 || j > 5 {
			break
		}
		if inc[i] > 0 && inc[j] > 0 {
			minij := min(inc[i], inc[j])
			inc[i] -= minij
			inc[j] -= minij
			if i == 0 && j == 3 {
				A[0] += minij
			} else if i == 0 && j == 4 {
				A[1] += minij
			} else if i == 0 && j == 5 {
				A[2] += minij
			} else if i == 1 && j == 3 {
				A[3] += minij
			} else if i == 1 && j == 4 {
				A[4] += minij
			} else if i == 1 && j == 5 {
				A[5] += minij
			} else if i == 2 && j == 3 {
				A[6] += minij
			} else if i == 2 && j == 4 {
				A[7] += minij
			} else if i == 2 && j == 5 {
				A[8] += minij
			}
		}
		if inc[i] == 0 && inc[j] > 0 {
			i++
		}
		if inc[i] > 0 && inc[j] == 0 {
			j++
		}
		if inc[i] == 0 && inc[j] == 0 {
			i++
			j++
		}
	}

	return A
}

func max(arr []int) int {
	m := 0
	for _, a := range arr {
		if a > m {
			m = a
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValid(inc []int) bool {
	return len(inc) == 6 && inc[0]+inc[1]+inc[2] == inc[3]+inc[4]+inc[5]
}
