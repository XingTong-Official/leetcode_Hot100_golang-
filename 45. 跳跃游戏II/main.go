package main

func main() {

}
func jump(nums []int) int {
	ans := 0
	local := 0
	for local < len(nums)-1 {
		ans++
		num := nums[local]
		if num+local >= len(nums)-1 {
			break
		}
		m := 0
		temp := local
		for i := 1; i <= num; i++ {
			if nums[local+i]+local+i >= m {
				m = nums[local+i] + local + i
				temp = local + i
			}
		}
		local = temp
	}
	return ans
}
