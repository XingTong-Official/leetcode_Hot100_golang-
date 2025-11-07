package main

import "fmt"

func main() {
	s := ""
	answer := lengthOfLongestSubstring(s)
	fmt.Println(answer)
}
func lengthOfLongestSubstring(s string) int {
	r := []rune(s)
	m := 0
	left, right := 0, 0
	for i, ch := range s {
		if i == 0 {
			m = 1
			continue
		}
		for j := left; j <= right; j++ {
			if r[j] == ch {
				left = j + 1
				break
			}
		}
		right = i
		m = max(m, right-left+1)
	}
	return m
}
