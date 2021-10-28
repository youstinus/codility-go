/*
Let us define a transformation that, given a string S of letters 'a' and/or 'b', replaces some consecutive sequence "abb" in S by "baa".

For example, after applying such a transformation to the string "abbabb", both strings "baaabb" and "abbbaa" can be obtained.

Write a function:

    func Solution(S string) string

that, given a string S of length N, returns the alphabetically largest string that can be obtained by performing the above operation any number of times.

Examples:

1. Given S = "ababb", your function should return "baaaa".
"baaaa" is alphabetically the largest possible outcome and may be obtained from "ababb" by using two transformations:
"ababb" → "abbaa" → "baaaa"

2. Given S = "abbbabb", your function should return "babaaaa".
"babaaaa" may be obtained from "abbbabb" by using three transformations:
"abbbabb" → "abbbbaa" → "baabbaa" → "babaaaa"

3. Given S = "aaabab", your function should return "aaabab".
No operation can be performed on S since it contains no "abb" fragment.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [1..100,000];
        string S consists only of the characters "a" and/or "b".

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.DeepEqual(Solution("ababb"), "baaaa"))
	fmt.Println(reflect.DeepEqual(Solution("abbbabb"), "babaaaa"))
	fmt.Println(reflect.DeepEqual(Solution("aaabab"), "aaabab"))
}
