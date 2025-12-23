package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	set := make(map[string]bool, len(wordDict))
	for _, s := range wordDict {
		set[s] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if set[s[j:i]] && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
