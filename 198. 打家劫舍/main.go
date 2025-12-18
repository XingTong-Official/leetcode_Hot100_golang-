package main

func main() {
	rob([]int{1, 2, 3, 1})
}
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		m := 0
		for j := 0; j < i-1; j++ {
			m = max(m, dp[j]+nums[i])
		}
		dp[i] = m
	}
	return max(dp[len(nums)-1], dp[len(nums)-2])
}
