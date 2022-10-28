package main

// 100/100/100
func Solution(A []int) int {
	mp := make(map[int]interface{})
	for _, a := range A {
		mp[a] = struct{}{}
	}

	return len(mp)
}
