package main

import (
	"fmt"
)

func main() {
	s1 := "abab"
	s2 := "ab"
	answer := findAnagrams(s1, s2)
	fmt.Println(answer)
}

// 优化移动窗口，不能用map，用数组替代，map性能远差于数组
func findAnagrams(s string, p string) []int {
	answer := []int{}
	if len(p) > len(s) {
		return answer
	}
	length := len(p)
	m := initMap(s, p)
	if calculateDifference(m) == 0 {
		answer = append(answer, 0)
	}
	for left := 1; left <= len(s)-length; left++ {
		m[int(s[left-1])-97]--
		m[int(s[left+length-1]-97)]++
		if calculateDifference(m) == 0 {
			answer = append(answer, left)
		}
	}
	return answer
}
func initMap(s string, p string) []int {
	length := len(p)
	m := make([]int, 26)
	for i := 0; i < length; i++ {
		m[int(s[i])-97]++
		m[int(p[i])-97]--
	}
	return m
}
func calculateDifference(m []int) int {
	difference := 0
	for i := 0; i < 26; i++ {
		t := m[i]
		if t < 0 {
			t = -t
		}
		difference += t
	}
	return difference
}

//方法很差，时间复杂度相同，但是很差，独立每次计算了窗口内所有内容
// func findAnagrams(s string, p string) []int {
// 	pf := []byte(p)
// 	nums := make([]byte, 26)
// 	nums2 := make([]byte, 26)
// 	for _, b := range pf {
// 		nums[b-97]++
// 	}
// 	answer := []int{}
// 	length := len(p)
// 	for left := 0; left <= len(s)-length; left++ {
// 		sub := s[left : left+length]
// 		difference := isAnagrams(sub, nums, nums2)
// 		if difference == 0 {
// 			answer = append(answer, left)
// 		} else {
// 			left += (difference / 2) - 1
// 		}
// 	}
// 	return answer
// }

// // 计算字符差值数量
// func isAnagrams(s string, p []byte, nums []byte) int {
// 	sb := []byte(s)
// 	for _, b := range sb {
// 		nums[b-97]++
// 	}
// 	totalDifference := 0
// 	for i := 0; i < 26; i++ {
// 		difference := 0
// 		if nums[i] >= p[i] {
// 			difference = int(nums[i] - p[i])
// 		} else {
// 			difference = int(p[i] - nums[i])
// 		}
// 		totalDifference += int(difference)
// 		nums[i] = 0
// 	}
// 	return totalDifference
// }

//排序nlogn
// func isAnagrams(s string, p []byte) bool {
// 	sf := []byte(s)
// 	sort.Slice(sf, func(i, j int) bool { return sf[i] < sf[j] })
// 	return bytes.Equal(sf, p)
// }
