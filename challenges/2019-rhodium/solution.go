package main

// Solution exported
// add functionality to check cities that you will be someday
func Solution(T []int) int {

	n := len(T)

	sum := 0

	for k := 0; k < n-1; k++ {
		was := make([]int, n+1)
		//be := make([]int, n+1)
		was[k] = 1
		for i := k + 1; i < n; i++ {
			check := false
			for j, a := range T {
				if (i == a && was[j] == 1) || (was[a] == 1 && i == j) {
					check = true
					break
				}
			}

			if check {
				sum++
				was[i] = 1
			} else {
				break
			}
		}
	}

	return sum + n
}
