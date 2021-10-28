package main

// 100/100/100
func Solution(A []int) int {
	// write your code in Go 1.4
	// variables int64 used because it can handle worst case scenario 100,000*1,000,000,000
	var devide int64 = 1000000007
	var min int64 = 9223372036854775807 // max int64
	var totalSum int64                  // saves total sum before highway
	eastern := A[len(A)-1]              // gets eastern most city value

	// sums all city values without highway
	for _, a := range A {
		totalSum += int64(eastern - a)
	}

	// checks cities one by one with highway
	// for every city to the left of selected one, the travel time decreases by highway value
	for i, a := range A {
		// formula to get total value with one highway
		sum := totalSum - int64(i+1)*int64(eastern-a)
		// checks if found shorter value is smaller than already calculated
		if sum < min {
			min = sum
		} /* else { // this case is false. Do not use, returns too early sometimes
			// This case returned 38/42/33 performance might be a problem when dealing with int64
			// if value is rising, then must return because value goes down
			// and then increases to the last city
			return int(min % devide)
		}*/
	}

	// gets remainder and transforms to int32 value
	return int(min % devide)
}
