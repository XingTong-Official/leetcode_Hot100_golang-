package main

func main() {
	coinChange([]int{2}, 3)
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		mn := amount + 1
		for _, value := range coins {
			if value > i || dp[i-value] == -1 {
				continue
			}
			mn = min(mn, dp[i-value]+1)
		}
		if mn == amount+1 {
			dp[i] = -1
		} else {
			dp[i] = mn
		}

	}
	return dp[amount]
}

/**

dp转移方程:
dp[n]=min( dp[n-coins[0]]+1,dp[n-coins[1]]+1...)

*/
