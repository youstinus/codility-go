package main

// 100/100/100
func Solution(A []int) int {
	length := len(A)
	sum := 0
	mp := make(map[int]int)
	max := 0
	for _, a := range A {
		sum += a
		_, ok := mp[a]
		if ok {
			return 0
		}
		mp[a]++
		if a > max {
			max = a
		}
	}

	if (length)*(1+length)/2 != sum || max != length {
		return 0
	}

	return 1
}
