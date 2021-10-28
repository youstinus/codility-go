package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Solution exported
func Solution(A [][]int) int {
	flips := processColumns(A)

	max := 0

	wg := sync.WaitGroup{}
	c := make(chan int)
	for _, v := range flips {
		wg.Add(1)
		go countScore(A, v, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		if v > max {
			max = v
		}
	}

	return max
}

func processColumns(A [][]int) [][]int8 {
	n := len(A[0])
	n2 := n * n
	flips := make([][]int8, 0)
	for i := 0; i < n2; i++ {
		s := strings.Split(strconv.FormatInt(int64(i), 2), "")
		choose := make([]int8, n)
		for j, v := range s {
			v1, err := strconv.Atoi(v)
			if err == nil {
				choose[j] = int8(v1)
			}
		}

		flips = append(flips, choose)
	}

	return flips
}

func countScore(A [][]int, flips []int8, cc chan<- int, wwg *sync.WaitGroup) {
	sum := 0
	c := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < len(A); i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			tmpWh := (A[ii][0] + int(flips[0])) % 2
			for i2, num := range A[ii] {
				if (num+int(flips[i2]))%2 != tmpWh {
					c <- 0
					return
				}
			}
			c <- 1
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		sum += v
	}

	wwg.Done()

	cc <- sum
}

// Solution2 exported second solution
func Solution2(A [][]int) int {
	b := whichChanges(A)
	fmt.Println(b)
	for i := 0; i < len(A); i++ {

	}
	return 0
}

func whichChanges(A [][]int) [][]int {
	indexes := make([][]int, len(A))
	for i1, v1 := range A {
		c0 := make([]int, 0)
		c1 := make([]int, 0)
		for i2, v2 := range v1 {
			if v2 == 0 {
				c0 = append(c0, i2)
			} else {
				c1 = append(c1, i2)
			}
		}

		if len(c0) < len(c1) {
			indexes[i1] = c0
		} else {
			indexes[i1] = c1
		}
	}
	return indexes
}
