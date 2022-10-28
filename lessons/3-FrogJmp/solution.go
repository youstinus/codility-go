package main

// 100/100/100
func Solution(X int, Y int, D int) int {
	x := (Y - X) / D
	if (Y-X)%D > 0 {
		x++
	}
	return x
}
