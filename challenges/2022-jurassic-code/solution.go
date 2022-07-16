package main

// Solution 100/100/100
func Solution(X []int, Y []int, colors string) int {
	cols := []rune(colors)
	histo := make(map[int][]*Point)
	keys := make([]int, 0)
	maxi := 0
	for i := 0; i < len(X); i++ {
		num := X[i]*X[i] + Y[i]*Y[i]

		_, ok := histo[num]
		if !ok {
			keys = append(keys, num)
		}

		histo[num] = append(histo[num],
			&Point{
				Length: num,
				Color:  cols[i] == 82,
				Index:  i,
			},
		)
		if num > maxi {
			maxi = num
		}
	}

	keys = MergeSort(keys)

	red, green := 0, 0
	max := 0
	for _, i := range keys {
		his, ok := histo[i]
		if !ok {
			continue
		}
		for _, a := range his {
			if a.Color { // R
				red++
			} else { // G
				green++
			}
		}

		if red == green && max < red {
			max = red
		}
	}

	return max * 2
}

type Point struct {
	Length int
	Color  bool
	Index  int
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := MergeSort(data[:middle])
	right := MergeSort(data[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}
