package main

func main() {
	permute([]int{1, 2})
}
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{append([]int{}, nums...)}
	}
	answer := [][]int{}
	for index, num := range nums {
		rest := append([]int{}, nums[:index]...) // copy left
		rest = append(rest, nums[index+1:]...)   // then append right
		arrays := permute(rest)

		for _, array := range arrays {
			answer = append(answer, append([]int{num}, array...))
		}
	}
	return answer
}
