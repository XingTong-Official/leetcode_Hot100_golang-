package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{0, 0, 0, 0}
	answer := threeSum(nums)
	fmt.Println(answer)
}

// 双指针法
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	answer := [][]int{}
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		if i == len(nums)-1 {
			break
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if num+nums[left]+nums[right] == 0 {
				answer = append(answer, []int{num, nums[left], nums[right]})
				left++
				for left < len(nums)-1 && nums[left-1] == nums[left] {
					left++
				}
			} else if num+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return answer
}

// // 哈希方法，很困难,O(n²)复杂度，但是效率很差
// func threeSum(nums []int) [][]int {
// 	first := map[int]bool{} //第一个数字不能重复

// 	ex := map[string]bool{} //检查一个三元数字是否重复
// 	answer := [][]int{}     //结果数组
// 	m := map[int]int{}      //检查是否存在对应结果
// 	for i, num := range nums {
// 		m[-num] = i
// 	}
// 	p1, p2 := 0, 0
// 	//第一个数字遍历
// 	for index, num1 := range nums {
// 		//防止第一个数字重复计算
// 		if first[num1] {
// 			continue
// 		}
// 		first[num1] = true
// 		p1 = index
// 		second := map[int]bool{}
// 		//第二个数字遍历
// 		for p2 = p1 + 1; p2 < len(nums); p2++ {
// 			//防止第二个数字重复计算
// 			if second[nums[p2]] {
// 				continue
// 			}
// 			second[nums[p2]] = true
// 			p3, ok := m[nums[p2]+nums[p1]]
// 			if ok && p1 < p2 && p2 < p3 { //要求指针的顺序，确保数组唯一性，但是仍可能存在不同数组，如[0,1.-1],[1,0,-1]
// 				arr1 := []int{nums[p1], nums[p2], -nums[p2] - nums[p1]}
// 				//二次去重，[0,1.-1],[1,0,-1]确保这样的数组不会重复进入answer,排序后作为key
// 				sort.Ints(arr1)
// 				key := fmt.Sprintf("%v", arr1)
// 				_, ok := ex[key]
// 				if ok {
// 					continue
// 				} else {
// 					ex[key] = true
// 					answer = append(answer, arr1)
// 				}
// 			}
// 		}
// 	}
// 	return answer
// }
