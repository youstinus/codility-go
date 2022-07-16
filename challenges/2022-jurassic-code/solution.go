package main

// Solution 23/33/9
func Solution(X []int, Y []int, colors string) int {
	cols := []rune(colors)
	histo := make([]*Point, len(X))
	color := true
	for i := 0; i < len(X); i++ {
		num := X[i]*X[i] + Y[i]*Y[i]
		if cols[i] == 82 { // R
			color = true
		} else { // G
			color = false
		}

		histo[i] = &Point{
			Length: num,
			Color:  color,
			Index:  i,
		}
	}

	BubbleSort(histo)

	red, green := 0, 0
	max := 0
	for _, h := range histo {
		if h.Color { // R
			red++
		} else { // G
			green++
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

func Sort(x []*Point) []*Point {
	// for i
	return nil
}

func BubbleSort(data []*Point) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j].Length < data[j-1].Length {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}
