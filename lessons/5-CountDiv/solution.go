package main

// 100/100/100
func Solution(A int, B int, K int) int {
	mod := A % K
	start := A
	if mod > 0 {
		start += K - mod
	}

	if start > B {
		return 0
	}

	dali := ((B - start) / K) + 1

	return dali
}

// 50/100/0
func Solution2(A int, B int, K int) int {
	count := 0
	start := A % K
	end := B - A + start

	for i := start; i <= end; i++ {
		if i%K == 0 {
			count++
		}
	}

	return count
}

// 50/100/0
func Solution1(A int, B int, K int) int {
	count := 0
	for i := A; i <= B; i++ {
		if i%K == 0 {
			count++
		}
	}

	return count
}
