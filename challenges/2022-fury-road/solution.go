package main

import "fmt"

// 100/100/100
func Solution(R string) int {
	if len(R) == 0 {
		return 0
	}

	length := len(R)
	scooter := make([]int, length)
	footer := make([]int, length)
	prevScoot := 0
	prevFoot := 0
	addScoot := 0
	addFoot := 0

	for i, r := range R {
		if r == 'A' {
			addScoot = 5
		} else if r == 'S' {
			addScoot = 40
		} else {
			panic(r)
		}

		if R[length-i-1] == 'A' {
			addFoot = 20
		} else if R[length-i-1] == 'S' {
			addFoot = 30
		} else {
			panic(r)
		}

		scooter[i] = prevScoot + addScoot
		footer[length-i-1] = prevFoot + addFoot
		prevScoot = scooter[i]
		prevFoot = footer[length-i-1]
	}

	var min int = footer[0]
	index := -1

	for i := 0; i < length-1; i++ {
		sum := scooter[i] + footer[i+1]
		if sum < min {
			min = sum
			index = i
		}
	}

	if scooter[length-1] < min {
		min = scooter[length-1]
		index = length
	}

	fmt.Println(min, index)
	return min
}
