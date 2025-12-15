package main

func main() {

}
func partitionLabels(s string) []int {
	ans := []int{}
	m := make(map[rune]int, 26)
	for i, r := range s {
		if _, ok := m[r]; !ok {
			m[r] = i
		} else {

		}

	}
}
