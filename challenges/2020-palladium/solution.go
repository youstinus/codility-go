package main

// Solution2 exported
// 25-50-0 but 220 line maximum golden reward 100-100-100
func Solution2(H []int) int {
	nH, maxIndexes := maxes(H, max(H))
	subLeft := nH[:maxIndexes[0]]
	maxSubLeft := max(subLeft)
	subRight := nH[maxIndexes[len(maxIndexes)-1]+1:]
	maxSubRight := max(subRight) + maxIndexes[len(maxIndexes)-1] + 1

	if maxSubRight >= len(H) {
		maxSubRight = maxIndexes[0]
	}

	sum0 := H[maxIndexes[0]] * (len(H) - maxIndexes[0])
	sum1 := H[maxIndexes[0]] * (maxIndexes[len(maxIndexes)-1] + 1)
	sum2 := H[maxSubLeft] * maxIndexes[0]
	sum3 := H[maxSubRight] * (len(H) - maxIndexes[len(maxIndexes)-1] - 1)

	sumLeft := sum0 + sum2
	sumRight := sum1 + sum3
	if sumLeft < sumRight {
		return sumLeft
	}
	return sumRight
}

// Solution1 exported 58-100-16
func Solution1(H []int) int {
	minSum := 2147483647
	for i := range H {
		max1 := max01(H[:i])
		max2 := max01(H[i:]) + i
		sum := 0
		if max1 != -1 {
			sum += H[max1] * i
		}
		if max2 != -1 {
			sum += H[max2] * (len(H) - i)
		}
		if minSum > sum {
			minSum = sum
		}
	}
	return minSum
}

func max01(H []int) int {
	ind := -1
	max := -1
	for i, a := range H {
		if a > max {
			max = a
			ind = i
		}
	}
	return ind
}

func maxes(H []int, maxInd int) ([]int, []int) {
	indexes := make([]int, 0)
	nH := make([]int, len(H))
	for i, a := range H {
		nH[i] = H[i]
		if a == H[maxInd] {
			indexes = append(indexes, i)
			nH[i] = 0
		}
	}
	return nH, indexes
}

// Solution3 exported 25-50-0
func Solution3(H []int) int {
	firstInd, lastInd := maxIndexes2(H)
	return twoIndexes(H, firstInd, lastInd)
}

func maxIndexes2(H []int) (int, int) {
	first := 0
	last := 0
	ma := 0
	for i := range H {
		if H[i] == ma {
			last = i
		}
		if H[i] > ma {
			ma = H[i]
			first = i
			last = i
		}
	}
	return first, last
}

func maxIndex2(H []int) int {
	ind := 0
	max := 0
	for i, a := range H {
		if a > max {
			max = a
			ind = i
		}
	}
	return ind
}

func twoIndexes(H []int, firstInd int, lastInd int) int {
	sum0 := H[lastInd] * (lastInd + 1)
	sum1 := H[firstInd] * (len(H) - firstInd)
	sum2 := 0
	if firstInd > 0 {
		sum2 = (H[maxIndex2(H[:firstInd])]) * firstInd
	}
	sum3 := 0
	if lastInd < len(H)-1 {
		sum3 = (H[maxIndex2(H[lastInd+1:])+lastInd+1]) * (len(H) - lastInd - 1)
	}
	return minimum2(sum0+sum3, sum1+sum2)
}

func minimum2(first int, second int) int {
	if first > second {
		return second
	}
	return first
}

// Solution4 exported 33-50-16
func Solution4(H []int) int {
	n := len(H)
	maxesFromLeft, maxesFromRight := getMaxes3(H)
	firstInd, lastInd := maxIndexes3(H)
	sumRight, sumLeft := someIndexes3(H, firstInd, lastInd)
	for i := lastInd + 1; i < n-1; i++ {
		newSumRight := sumRight + H[lastInd] - maxesFromRight[i]*(n-i+1) + maxesFromRight[i+1]*(n-i+2)
		if newSumRight < sumRight {
			sumRight = newSumRight
		}
	}
	for i := firstInd - 1; i > 0; i-- {
		newSumLeft := sumLeft + H[firstInd] - maxesFromLeft[i]*(i+1) + maxesFromLeft[i-1]*i
		if newSumLeft < sumLeft {
			sumLeft = newSumLeft
		}
	}
	return minimum2(sumRight, sumLeft)
}

func getMaxes3(H []int) ([]int, []int) {
	n := len(H)
	maxesFromLeft := make([]int, len(H))
	maxLeft := 0
	maxesFromRight := make([]int, len(H))
	maxRight := 0
	for i := range H {
		if maxLeft < H[i] {
			maxLeft = H[i]
		}
		if maxRight < H[n-1-i] {
			maxRight = H[n-1-i]
		}
		maxesFromLeft[i] = maxLeft
		maxesFromRight[n-1-i] = maxRight
	}
	return maxesFromLeft, maxesFromRight
}

func someIndexes3(H []int, firstInd int, lastInd int) (int, int) {
	sum0 := H[lastInd] * (lastInd + 1)
	sum1 := H[firstInd] * (len(H) - firstInd)
	sum2 := 0
	if firstInd > 0 {
		sum2 = (H[maxIndex3(H[:firstInd])]) * firstInd
	}
	sum3 := 0
	if lastInd < len(H)-1 {
		sum3 = (H[maxIndex3(H[lastInd+1:])+lastInd+1]) * (len(H) - lastInd - 1)
	}
	return sum0 + sum3, sum1 + sum2
}

func maxIndexes3(H []int) (int, int) {
	first := 0
	last := 0
	ma := 0
	for i := range H {
		if H[i] == ma {
			last = i
		}
		if H[i] > ma {
			ma = H[i]
			first = i
			last = i
		}
	}
	return first, last
}

func maxIndex3(H []int) int {
	ind := 0
	max := 0
	for i, a := range H {
		if a > max {
			max = a
			ind = i
		}
	}
	return ind
}

func minimum3(first int, second int) int {
	if first > second {
		return second
	}
	return first
}

// Solution exported 58-100-16 100-100-100
func Solution(H []int) int {
	maxesFromLeft, maxesFromRight := getMaxes(H)
	minSum := 2147483647
	for i := range H {
		max1 := maxesFromLeft[i]
		max2 := maxesFromRight[i]
		sum := max1*(i+1) + max2*(len(H)-i) - maximum(max1,max2)
		if minSum > sum {
			minSum = sum
		}
	}
	return minSum
}

func getMaxes(H []int) ([]int, []int) {
	n := len(H)
	maxesFromLeft := make([]int, len(H))
	maxLeft := 0
	maxesFromRight := make([]int, len(H))
	maxRight := 0
	for i := range H {
		if maxLeft < H[i] {
			maxLeft = H[i]
		}
		if maxRight < H[n-1-i] {
			maxRight = H[n-1-i]
		}
		maxesFromLeft[i] = maxLeft
		maxesFromRight[n-1-i] = maxRight
	}
	return maxesFromLeft, maxesFromRight
}

func max(H []int) int {
	ind := -1
	max := -1
	for i, a := range H {
		if a > max {
			max = a
			ind = i
		}
	}
	return ind
}

func maximum(first int, second int) int {
	if first < second {
		return second
	}
	return first
}