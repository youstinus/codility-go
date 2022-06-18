package main

// 28/50/12
func Solution(A []int, X []int) int {
	max := 0
	for i := range A {
		tmp := went(i, 0, A, X, make([]int, len(A)))
		if tmp == len(A) {
			return tmp
		}

		if tmp > max {
			max = tmp
		}
	}

	return max
}

func went(index int, fuel int, A, X, was []int) int {
	was[index] = 1
	fuel += A[index]
	max := 0
	for i := range A {
		if index == i || was[i] != 0 {
			continue
		}

		diff := Abs(X[i] - X[index])
		if diff > fuel {
			continue
		}

		fuel -= diff
		dst := make([]int, len(was))
		copy(dst, was)
		tmp := went(i, fuel, A, X, dst)
		if tmp > max {
			max = tmp
		}
	}

	return 1 + max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
