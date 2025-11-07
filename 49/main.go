package main

import "fmt"

func main() {
	test := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	answer := groupAnagrams(test)
	fmt.Println(answer)
}
func groupAnagrams(strs []string) [][]string {
	answer := [][]string{}
	m := make(map[string][]string)
	for _, str := range strs {
		s := calcuString(str)
		m[s] = append(m[s], str)
	}
	for _, strss := range m {
		answer = append(answer, strss)
	}
	return answer
}

//计算字符串特征值
func calcuString(str string) string {
	s := []rune(str)
	ints := [26]int{}
	for i := 0; i < len(s); i++ {
		s1 := s[i] - 97
		ints[s1] = ints[s1] + 1
	}
	s2 := "a" + string(ints[0]) + "b" + string(ints[1]) + "c" + string(ints[2]) + "d" + string(ints[3]) + "e" + string(ints[4]) + "f" + string(ints[5]) + "g" + string(ints[6]) + "h" + string(ints[7]) + "i" + string(ints[8]) + "j" + string(ints[9]) + "k" + string(ints[10]) + "l" + string(ints[11]) + "m" + string(ints[12]) + "n" + string(ints[13]) + "o" + string(ints[14]) + "p" + string(ints[15]) + "q" + string(ints[16]) + "r" + string(ints[17]) + "s" + string(ints[18]) + "t" + string(ints[19]) + "u" + string(ints[20]) + "v" + string(ints[21]) + "w" + string(ints[22]) + "x" + string(ints[23]) + "y" + string(ints[24]) + "z" + string(ints[25])
	return s2
}
