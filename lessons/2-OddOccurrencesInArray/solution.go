package main

// 100/100/100
func Solution(A []int) int {
	odd := make(map[int]int)
	for _, a := range A {
		odd[a] += 1
		if odd[a] == 2 {
			delete(odd, a)
		}
	}

	for k := range odd {
		return k
	}

	return 0
}
