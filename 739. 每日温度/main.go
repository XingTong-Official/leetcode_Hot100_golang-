package main

func main() {
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
}
func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := []int{}
	for index, t := range temperatures {

		for len(stack) != 0 && t > temperatures[stack[len(stack)-1]] {
			now := stack[len(stack)-1]
			answer[now] = index - now
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, index)
	}
	return answer
}
