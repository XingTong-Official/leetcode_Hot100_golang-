package main

import "fmt"

func main() {
	nums := []int{2, 5, 3, 1, 5, 1}
	answer := findRepeatDocument(nums)
	fmt.Println(answer)
}
func findRepeatDocument(documents []int) int {

	// 循环怼
	n := len(documents)
	i := 0
	for i < n {
		x := documents[i]
		if x == i {
			i++
			continue
		}
		if documents[x] == x {
			return x
		} else {
			documents[x], documents[i] = documents[i], documents[x]
		}
	}
	return -1
}

// func findRepeatDocument(documents []int) int {
// 	for index, num := range documents {
// 		documents[index], documents[num] = documents[num], documents[index]
// 	}
// 	for index, num := range documents {
// 		documents[index], documents[num] = documents[num], documents[index]
// 	}
// 	for index, num := range documents {
// 		if num != index {
// 			return num
// 		}
// 	}
// 	return 0
// }
