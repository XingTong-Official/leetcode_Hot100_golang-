package main

//n²复杂度,效率极差
// func moveZeroes(nums []int) {
// 	p1 := -1
// 	for index, _ := range nums {
// 		for i := 0; i < len(nums); i++ {
// 			if nums[i] == 0 {
// 				p1 = i
// 				break
// 			}
// 		}
// 		if p1 >= index {
// 			continue
// 		}
// 		nums[index], nums[p1] = nums[p1], nums[index]
// 	}
// }
import "fmt"

func main() {
	test := []int{1, 0, 1}
	moveZeroes(test)
	fmt.Println(test)
}

//双指针,p1指向0的数字，p2指向非零数字
// func moveZeroes(nums []int) {
// 	p1, p2 := 0, 0
// 	for {
// 		if p1 == len(nums) || p2 == len(nums) {
// 			break
// 		}
// 		if nums[p1] == 0 {
// 			if nums[p2] != 0 {
// 				if p1 < p2 {
// 					nums[p2], nums[p1] = nums[p1], nums[p2]
// 				}
// 				p2++
// 			} else {
// 				p2++
// 			}
// 		} else {
// 			p1++
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	left, right, length := 0, 0, len(nums)
	for right < length {
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
		right++
	}
}
