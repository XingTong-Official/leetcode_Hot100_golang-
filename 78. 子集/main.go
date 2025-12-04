package main

func main() {
	subsets([]int{3, 2, 4, 1})
}
func subsets(nums []int) [][]int {
	answer := [][]int{}
	var subset func(array []int, index int)
	subset = func(array []int, index int) {
		if index == len(nums) {
			answer = append(answer, array)
			return
		}
		arr := make([]int, len(array))
		copy(arr, array)
		subset(arr, index+1)

		arr = append(arr, nums[index])
		subset(arr, index+1)
	}
	subset([]int{}, 0)
	return answer
}
