package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
func combinationSum(candidates []int, target int) [][]int {
	answer := [][]int{}
	var combination func(path []int, num int, index int)
	combination = func(path []int, num int, index int) {
		if num == target {
			answer = append(answer, append([]int{}, path...))
			return
		}
		if index == len(candidates) || num > target {
			return
		}
		path = append(path, candidates[index])
		combination(path, num+candidates[index], index)
		path = path[:len(path)-1]
		combination(path, num, index+1)

	}
	combination([]int{}, 0, 0)
	return answer
}

// func combinationSum(candidates []int, target int) [][]int {
// 	answer := [][]int{}
// 	type pathStruct struct {
// 		p   []int
// 		sum int
// 	}
// 	path := pathStruct{
// 		p:   []int{},
// 		sum: 0,
// 	}
// 	var combination func(path pathStruct) bool
// 	combination = func(path pathStruct) bool {

// 		if path.sum == target {
// 			sort.Slice(path.p, func(i, j int) bool { return path.p[i] < path.p[j] })
// 			for _, nums := range answer {
// 				if slices.Equal(nums, path.p) {
// 					return false
// 				}
// 			}
// 			answer = append(answer, path.p)
// 			return false
// 		}
// 		if path.sum > target {
// 			return false
// 		}
// 		for _, num := range candidates {
// 			path.p = append(path.p, num)
// 			path.sum += num
// 			if !combination(pathStruct{
// 				p:   append([]int{}, path.p...),
// 				sum: path.sum,
// 			}) {
// 				break
// 			}
// 			path.p = path.p[:len(path.p)-1]
// 			path.sum -= num
// 		}
// 		return true
// 	}
// 	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
// 	combination(path)
// 	return answer
// }
