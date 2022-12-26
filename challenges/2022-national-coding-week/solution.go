package main

// 29/45/0
func Solution(A []int, B []int) int {
	// println("prin")
	// var tree Node

	nodes := make([]*Node, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		nodes[i] = &Node{
			I:     i,
			Links: make([]*Node, 0),
		}
	}

	for i := range A {
		a, b := A[i], B[i]
		nodes[a].Links = append(nodes[a].Links, nodes[b])
		nodes[b].Links = append(nodes[b].Links, nodes[a])
	}

	//remove
	// start, length := biggestChain(nodes)
	// println("bigest chain start", start, "length", length)
	matrix = make(map[int]map[int]int)
	return count3(nodes)

	// return count(nodes)

	// return recurse(nodes[0], -1, 1)
}

var matrix = make(map[int]map[int]int)

func count3(nodes []*Node) int {
	for i := 0; i < len(nodes); i++ {
		find2(nodes[i], nodes[i], 0, 1)
	}

	sum := 0
	for _, subMatrix := range matrix {
		for _, count := range subMatrix {
			if count%2 == 1 {
				sum++
			}
		}
	}
	return sum / 2
}

func find2(start, current *Node, previous, sum int) {
	subMatrix, ok := matrix[start.I]
	if !ok {
		subMatrix = make(map[int]int)
		matrix[start.I] = subMatrix
	}

	for _, node := range current.Links {
		subMatrix[node.I] = sum
		if node.I == previous {
			continue
		}
		find2(start, node, current.I, sum+1)
	}
}

// Solution 29/45/0
func Solution2(A []int, B []int) int {
	println("prin")
	// var tree Node

	nodes := make([]*Node, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		nodes[i] = &Node{
			I:     i,
			Links: make([]*Node, 0),
		}
	}

	for i := range A {
		a, b := A[i], B[i]
		nodes[a].Links = append(nodes[a].Links, nodes[b])
		nodes[b].Links = append(nodes[b].Links, nodes[a])
	}

	//remove
	// start, length := biggestChain(nodes)
	// println("bigest chain start", start, "length", length)

	return count2(nodes)

	//return count(nodes)

	// return recurse(nodes[0], -1, 1)
}

func count2(nodes []*Node) int {
	sum := 0
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			length := find(nodes[i], nodes[j], 0, 0)
			if length%2 == 1 {
				sum++
			}
		}
	}
	return sum
}

func find(start, end *Node, previous, sum int) int { // length
	if start == end {
		return sum
	}

	for _, node := range start.Links {
		if node.I == previous {
			continue
		}
		length := find(node, end, start.I, sum+1)
		if length > 0 {
			return length
		}
	}

	return 0
}

func biggestChain(nodes []*Node) (int, int) {
	index, max := 0, 0

	for i, node := range nodes {
		length := maxLength(node, 0)
		if length > max {
			index = i
			max = length
		}
	}

	return index, max
}

func maxLength(node *Node, current int) int {
	if len(node.Links) == 0 {
		return current
	}

	max := 0
	for _, nod := range node.Links {
		num := maxLength(nod, current+1)
		if num > max {
			max = num
		}
	}

	return max
}

func recurse2(node *Node, index int) int {
	if len(node.Links) == 0 {
		return index
	}

	max := 0
	for _, nod := range node.Links {
		num := recurse2(nod, index+1)
		if num > max {
			max = num
		}
	}

	return max
}

func count(nodes []*Node) int {
	sum := 0

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {

		}
	}

	return sum
}

func recurse(node *Node, prev int, linkCount int) int {
	sum := 1
	for i := range node.Links {
		if node.Links[i].I == prev {
			continue
		} else if len(node.Links) == 1 && prev != -1 {

			return 1
		}

		innerSum := recurse(node.Links[i], node.I, linkCount+1)

		if linkCount%2 != 0 {
			sum += innerSum
		}

	}

	return sum
}

type Node struct {
	I     int // node number
	Links []*Node
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
