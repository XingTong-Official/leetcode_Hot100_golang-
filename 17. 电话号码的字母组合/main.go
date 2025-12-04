package main

func main() {
	letterCombinations("23")
}
func letterCombinations(digits string) []string {
	answer := []string{}
	path := []byte{}
	var letter func(path []byte, index int)
	letter = func(path []byte, index int) {
		if index == len(digits) {
			answer = append(answer, string(path))
			return
		}
		c := digits[index]
		if c == '2' {
			letter(append(append([]byte{}, path...), 'a'), index+1)
			letter(append(append([]byte{}, path...), 'b'), index+1)
			letter(append(append([]byte{}, path...), 'c'), index+1)
		} else if c == '3' {
			letter(append(append([]byte{}, path...), 'd'), index+1)
			letter(append(append([]byte{}, path...), 'e'), index+1)
			letter(append(append([]byte{}, path...), 'f'), index+1)
		} else if c == '4' {
			letter(append(append([]byte{}, path...), 'g'), index+1)
			letter(append(append([]byte{}, path...), 'h'), index+1)
			letter(append(append([]byte{}, path...), 'i'), index+1)
		} else if c == '5' {
			letter(append(append([]byte{}, path...), 'j'), index+1)
			letter(append(append([]byte{}, path...), 'k'), index+1)
			letter(append(append([]byte{}, path...), 'l'), index+1)
		} else if c == '6' {
			letter(append(append([]byte{}, path...), 'm'), index+1)
			letter(append(append([]byte{}, path...), 'n'), index+1)
			letter(append(append([]byte{}, path...), 'o'), index+1)
		} else if c == '7' {
			letter(append(append([]byte{}, path...), 'p'), index+1)
			letter(append(append([]byte{}, path...), 'q'), index+1)
			letter(append(append([]byte{}, path...), 'r'), index+1)
			letter(append(append([]byte{}, path...), 's'), index+1)
		} else if c == '8' {
			letter(append(append([]byte{}, path...), 't'), index+1)
			letter(append(append([]byte{}, path...), 'u'), index+1)
			letter(append(append([]byte{}, path...), 'v'), index+1)
		} else if c == '9' {
			letter(append(append([]byte{}, path...), 'w'), index+1)
			letter(append(append([]byte{}, path...), 'x'), index+1)
			letter(append(append([]byte{}, path...), 'y'), index+1)
			letter(append(append([]byte{}, path...), 'z'), index+1)
		}
	}
	letter(path, 0)
	return answer
}
