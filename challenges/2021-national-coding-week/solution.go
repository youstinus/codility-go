package main

import "strings"

// 63/100/20 timeout.
func Solution(S string) string {
	str := S
	for {
		ns := strings.Replace(str, "abb", "baa", -1)
		if ns != str {
			str = ns
		} else {
			break
		}
	}
	return str
}

// 63/100/20 timeout.
func Solution2(S string) string {
	str := []byte(S)
	changed := true
	for changed {
		changed = false
		for i := len(str) - 3; i >= 0; i-- {
			if str[i] == 'a' && str[i+1] == 'b' && str[i+2] == 'b' {
				changed = true
				str[i] = 'b'
				str[i+1] = 'a'
				str[i+2] = 'a'
			}
		}

	}
	return string(str)
}
