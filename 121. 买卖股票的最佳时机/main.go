package main

func main() {

}
func maxProfit(prices []int) int {
	answer := 0
	mi := prices[0]
	for _, num := range prices {
		answer = max(num-mi, answer)
		mi = min(mi, num)
	}
	return answer
}

// 单调栈
// func maxProfit(prices []int) int {
// 	stack := make([]int, 2000)
// 	answer := 0
// 	for i, num := range prices {
// 		for len(stack) != 0 && num <= prices[stack[len(stack)-1]] {
// 			stack = stack[:len(stack)-1]
// 		}

// 		stack = append(stack, i)
// 		price2 := prices[stack[len(stack)-1]] - prices[stack[0]]

// 		answer = max(price2, answer)
// 	}
// 	return answer
// }
