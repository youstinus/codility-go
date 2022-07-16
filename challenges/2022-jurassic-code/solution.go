package main

// Solution 53/66/36
func Solution(X []int, Y []int, colors string) int {
	cols := []rune(colors)
	histo := make(map[int][]*Point)
	color := true
	maxi := 0
	for i := 0; i < len(X); i++ {
		num := X[i]*X[i] + Y[i]*Y[i]
		if cols[i] == 82 { // R
			color = true
		} else { // G
			color = false
		}

		histo[num] = append(histo[num],
			&Point{
				Length: num,
				Color:  color,
				Index:  i,
			},
		)
		if num > maxi {
			maxi = num
		}
	}

	red, green := 0, 0
	max := 0
	for i := 1; i <= maxi; i++ {
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
