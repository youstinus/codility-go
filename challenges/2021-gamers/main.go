/*


You are given a string S of length N consisting of characters 'a' and 'b' with additional empty gaps represented by '?'. Your task is to replace every such gap with either an 'a' character or a 'b' character so that the longest fragment of S, consisting only of 'a' characters or 'b' characters, is as short as possible.

For example, for S = "aa??bbb", if you replace "??" with "aa", the longest fragment consisting of equal characters will have length 4: "aaaabbb". You can obtain a better result by replacing "??" with "ba", resulting in "aababbb". Then the longest fragment consisting of equal characters will have length 3.

Write a function:

    func Solution(S string) int

that, given a string S of length N, returns the minimum possible length of the longest fragment of S consisting of equal characters after replacing all "?" characters with letters.

Examples:

1. Given S = "aa??bbb", your function should return 3, as explained above.

2. Given S = "a?b?aa?b?a", your function should return 2. Question marks can be replaced in the following way: "aabbaabbaa".

3. Given S = "?????", your function should return 1. There are two possible ways to achieve this length: "ababa" and "babab".

4. Given S = "aa?b?aa", your function should return 3.

Write an efficient algorithm for the following assumptions:

        string S consists only of the following characters: "a", "b" and/or "?";
        N is an integer within the range [1..100,000].

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited. */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution("aa??bbb")) //3
	fmt.Println(Solution("aa?b?aa")) //2

	// fmt.Println(reflect.DeepEqual(Solution([]int{3, 7, 0, 5}), 9))
	// fmt.Println(reflect.DeepEqual(Solution([]int{1, 1, 1, 1, 1}), 1))
}
