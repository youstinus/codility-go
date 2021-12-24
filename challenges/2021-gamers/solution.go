package main

// 97 a
// 98 b
// 63 ?
func Solution(S string) int {
	r := []rune(S)
	var current rune = 0
	max := 0
	begin := 0
	for i, a := range r {
		if current == 63 || current != a {
			tmp := i - begin
			if tmp > max {
				max = tmp
			}
			begin = i
			current = a
		}
	}
	tmp := len(r) - begin
	if tmp > max {
		max = tmp
	}
	//return max

	ind1 := 0
	// ind2 := 0
	var prv rune = 0
	for i, a := range r {
		if a != prv {
			diff := i - ind1
			if diff > 0 {

			}
			ind1 = i
			prv = a
		}
	}
	// if a == 63 {
	// 	ind2 = i
	// } else {
	// 	diff := ind2 - ind1
	// 	if diff > 0 {

	// 	}
	// 	ind1 = i
	// }

	var prev rune
	//var next rune = []rune(S)[1]
	length := 0
	start := 0
	for i, a := range S {
		if a == 97 {
			if prev == 97 {

			} else {
				tmp := i - start
				if tmp > length {
					length = tmp
				}
				start = i
			}
			prev = 97
			continue
		}
		if a == 98 {
			if prev == 98 {

			} else {
				tmp := i - start
				if tmp > length {
					length = tmp
				}
				start = i
			}
			prev = 98
			continue
		}
		if a == 63 {
			//lete:=[]rune(S)[j+1]
			/*for j:=i+1;j<len(S);j++ {
				if []rune(S)[j] != a {
					lete :=
				}
			}*/
			if i > 0 {
				if prev == 97 {
					prev = 98
				} else if prev == 98 {
					prev = 97
				}
				tmp := i - start
				if tmp > length {
					length = tmp
				}
				start = i
			}
		}
	}
	/*tmp := len(S) - start
	if tmp > length {
		length = tmp
	}*/
	return length
}

// 97 a
// 98 b
// 63 ?
func Solution54(S string) int {
	var prev rune
	//var next rune = []rune(S)[1]
	length := 0
	start := 0
	for i, a := range S {
		if a == 97 {
			if prev == 97 {

			} else {
				tmp := i - start
				if tmp > length {
					length = tmp
				}
				start = i
			}
			prev = 97
			continue
		}
		if a == 98 {
			if prev == 98 {

			} else {
				tmp := i - start
				if tmp > length {
					length = tmp
				}
				start = i
			}
			prev = 98
			continue
		}
		if a == 63 {
			if i > 0 {
				if prev == 97 {
					prev = 98
				} else if prev == 98 {
					prev = 97
				}
				tmp := i - start
				if tmp > length {
					length = tmp
				}
				start = i
			}
		}
	}
	tmp := len(S) - start
	if tmp > length {
		length = tmp
	}
	return length
}

// 62/100/0
func Solution34(A []int) int {
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
