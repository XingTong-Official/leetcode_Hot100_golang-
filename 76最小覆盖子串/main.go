package main

import "fmt"

func main() {
	s1 := "a"
	s2 := "b"
	s := minWindow(s1, s2)
	fmt.Println(s)
}

//初始化存在极大问题，明天重构
func minWindow(s string, t string) string {
	if len(t) > len(s) || len(t) == 0 || len(s) == 0 {
		return ""
	}
	mt := map[uint8]int{}

	left, right := 0, 0
	finalAnswer := []uint8{}
	answer := []uint8{}
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}
	for right=0;right<len(s);right++{
			
	}


	for {
		right++
		if right == len(s) {
			return string(answer)
		}
		answer = append(answer, s[right])
		if _, ok := mt[s[right]]; ok {
			//如果右侧有有效值进入队列
			mt[s[right]]++
			for {
				leftValue := answer[left]
				answer = answer[left+1:]
				if _, ok := mt[leftValue]; ok {
					mt[leftValue]--
				}
				if calculateNums(mt) < 0 {
					answer = append([]uint8{leftValue}, answer...)
					mt[leftValue]++
					if len(answer) < len(finalAnswer) {
						finalAnswer = answer
					}
					break
				}
			}
		}
	}
	return string(finalAnswer)
}

func calculateNums(m map[uint8]int) int {
	nums := 0
	for _, num := range m {
		nums += num
		if num < 0 {
			return -1
		}
	}
	return nums
}
