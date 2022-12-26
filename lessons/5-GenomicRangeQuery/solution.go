package main

// 62/100/0
func Solution(S string, P []int, Q []int) []int {
	nums := make([]int, len(S))
	for i := range S {
		nums[i] = factor(S[i])
	}

	facts := make([]int, len(P))
	for i := range P {
		sub := nums[P[i] : Q[i]+1]
		facts[i] = minFactor(sub)
	}
	return facts
}

func minFactor(sub []int) int {
	min := 4
	for _, num := range sub {
		if num < min {
			min = num
			if min == 1 {
				return 1
			}
		}
	}

	return min
}

func factor(r byte) int {
	switch r {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	default:
		panic("wrong letter")
	}
}
