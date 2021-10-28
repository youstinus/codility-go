package main

// 62/100/0
func Solution(A []int) int {
	// write your code in Go 1.4
	hist := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		tmp := make([]int, len(A))
		copy(tmp, A)
		for j := len(tmp) - 1; j > i; j-- {
			if i < len(tmp)-1 {
				doub := tmp[j] / 2
				tmp[j-1] += doub
				tmp[j] -= (doub * 2)
			}
		}
		for k := 0; k < i; k++ {
			if i > 0 {
				dou := tmp[k] / 2
				tmp[k+1] += dou
				tmp[k] -= (dou * 2)
			}
		}
		hist[i] = tmp[i]
	}
	return max(hist)
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

// 62/100/0
func Solution2(A []int) int {
	m := 0
	tmp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		copy(tmp, A)

		if i < len(tmp)-1 {
			for j := len(tmp) - 1; j > i; j-- {
				if tmp[j] > 1 {
					doub := tmp[j] / 2
					tmp[j-1] += doub
					tmp[j] -= (doub * 2)
				}
			}
		}
		if i > 0 {
			for k := 0; k < i; k++ {
				if tmp[k] > 1 {
					dou := tmp[k] / 2
					tmp[k+1] += dou
					tmp[k] -= (dou * 2)
				}
			}
		}
		if tmp[i] > m {
			m = tmp[i]
		}
	}
	return m
}

// 25/40/0
func Solution3(A []int) int {
	maxNum := 0
	for i := range A {
		if A[i] > maxNum {
			maxNum = A[i]
		}
	}
	maxSliceInd := make([]int, 0)
	for i, a := range A {
		if a == maxNum {
			maxSliceInd = append(maxSliceInd, i)
		}
	}
	m := 0
	tmp := make([]int, len(A))
	for _, i := range maxSliceInd {
		copy(tmp, A)

		if i < len(tmp)-1 {
			for j := len(tmp) - 1; j > i; j-- {
				if tmp[j] > 1 {
					doub := tmp[j] / 2
					tmp[j-1] += doub
					tmp[j] -= (doub * 2)
				}
			}
		}
		if i > 0 {
			for k := 0; k < i; k++ {
				if tmp[k] > 1 {
					dou := tmp[k] / 2
					tmp[k+1] += dou
					tmp[k] -= (dou * 2)
				}
			}
		}
		if tmp[i] > m {
			m = tmp[i]
		}
	}
	return m
}
