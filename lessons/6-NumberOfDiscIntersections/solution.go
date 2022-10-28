package main

func Solution(A []int) int {
	sum := 0

	return sum
}

// 12/12/12
func Solution2(A []int) int {
	sum := 0
	mp := make(map[int]int)
	for i, a := range A {
		start := 0
		if i > a {
			start = i - a
		}

		end := len(A) - 1
		if len(A)-i-1 > a {
			end = i + a
		}

		for i := start; i <= end; i++ {
			mp[i] += 1
		}
	}

	for _, v := range mp {
		sum += (v / 2)
	}

	return sum
}

// 0/0/0
func Solution1(A []int) int {
	sum := 0

	for i, a := range A {
		if i < a {
			sum += i
		} else {
			sum += a
		}

		if len(A)-i-1 < a {
			sum += (len(A) - i - 1)
		} else {
			sum += a
		}
	}

	return sum
}
