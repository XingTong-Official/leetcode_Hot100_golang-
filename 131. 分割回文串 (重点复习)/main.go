package main

func main() {
	partition("aa")
}
func partition(s string) [][]string {
	answer := [][]string{}
	path := []string{}
	var part func(pre, local int)
	part = func(pre, local int) {
		if local == len(s)+1 {
			if pre == len(s) {
				answer = append(answer, append([]string{}, path...))
			}
			return
		}
		if isPalindrome(s[pre:local]) {
			path = append(path, s[pre:local])
			part(local, local+1)
			path = path[:len(path)-1]

		}
		part(pre, local+1)

	}
	part(0, 1)
	return answer
}
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
