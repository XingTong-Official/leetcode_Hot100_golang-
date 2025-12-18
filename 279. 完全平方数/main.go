package main

import "math"

func main() {
	numSquares(12)
}
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		mn := i
		jMax := int(math.Sqrt(float64(i)))
		for j := 1; j <= jMax; j++ {
			mn = min(mn, dp[i-j*j]+1)
		}
		dp[i] = mn
	}
	return dp[n]
}

/*
	dp[n]=min(range dp[n-1],dp[n-4]...+1)
*/
