package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}
func generateParenthesis(n int) []string {

	answer := []string{}
	path := make([]byte, 2*n)
	var generate func(n int, rest int)
	generate = func(left, right int) {
		if right == n {
			answer = append(answer, string(path))
			return
		}
		//填左括号
		if left < n {
			path[left+right] = '('
			generate(left+1, right)
		}
		if right < left {
			path[left+right] = ')'
			generate(left, right+1)
		}
	}
	generate(0, 0)
	return answer
}
