package main

func main() {

}
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		m := 1
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				m = max(m, dp[j]+1)
			}
		}
		dp[i] = m
	}
	m := 0
	for _, num := range dp {
		m = max(m, num)
	}
	return m
}
