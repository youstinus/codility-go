package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var N = 0

var M = 0

// Solution finds solution
func Solution(A [][]int) string {
	N = len(A)
	M = len(A[0])

	fmt.Println()
	start := time.Now()
	log.Println(start.Nanosecond())
	str := proceed(A, 0, 0)
	//time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Println("proceed", str)
	log.Println("took", end.Sub(start).Nanoseconds())

	fmt.Println()
	start = time.Now()
	log.Println(start.Nanosecond())
	str1 := proc(A)
	end = time.Now()
	fmt.Println("   proc", str1)
	log.Println("took", end.Sub(start).Nanoseconds())

	return str
}

// 77-100-33 timeout error > 7s
func proceed(A [][]int, row int, col int) string {

	if row+1 < N {
		if col+1 < M {
			if A[row][col+1] == A[row+1][col] {
				x1 := proceed(A, row+1, col)
				x2 := proceed(A, row, col+1)
				if x1 > x2 {
					return strconv.Itoa(A[row][col]) + x1
				} else {
					return strconv.Itoa(A[row][col]) + x2
				}
			} else if A[row][col+1] > A[row+1][col] {
				return strconv.Itoa(A[row][col]) + proceed(A, row, col+1)
			} else if A[row][col+1] < A[row+1][col] {
				return strconv.Itoa(A[row][col]) + proceed(A, row+1, col)
			}
		} else {
			return strconv.Itoa(A[row][col]) + proceed(A, row+1, col)
		}
	} else {
		if col+1 < M {
			return strconv.Itoa(A[row][col]) + proceed(A, row, col+1)
		} else {
			return strconv.Itoa(A[row][col])
		}
	}

	return ""
}

// 77-100-33 + runtime error exit with code 2
func proc(A [][]int) string {

	var mat = make([][]string, N)
	for i := range mat {
		mat[i] = make([]string, M+1)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			mat[i][j] = max(mat[i-1][j], mat[i][j-1]) + strconv.Itoa(A[i-1][j-1])
		}
	}

	return mat[N][M]
}

func max(x, y string) string {
	if x > y {
		return x
	}
	return y
}
