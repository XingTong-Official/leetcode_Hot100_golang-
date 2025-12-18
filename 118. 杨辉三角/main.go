package main

func main() {
	generate(5)
}
func generate(numRows int) [][]int {
	dp := [][]int{}
	for i := 0; i < numRows; i++ {
		dp = append(dp, []int{})
		for j := 0; j < i+1; j++ {
			if j == i || j == 0 {
				dp[i] = append(dp[i], 1)
				continue
			}
			local := dp[i-1][j] + dp[i-1][j-1]
			dp[i] = append(dp[i], local)
		}
	}
	return dp
}
