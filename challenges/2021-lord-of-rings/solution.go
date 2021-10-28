package main

// Solution exported
func Solution(S string, K int) string {
	res := []rune(S)
	count := 0
	for i := 0; i < K; i++ {
		for j := 0; j < len(res)-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
				count++
				break
			}
			if j >= len(res)-1 {
				return string(res)
			}
		}
	}
	return string(res)
}
