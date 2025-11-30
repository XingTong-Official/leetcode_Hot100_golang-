package main

func main() {
	decodeString("3[a2[c]]")
}
func decodeString(s string) string {
	stack := []rune{}
	for _, r := range s {
		if r == ']' {
			s := []rune{}
			for {
				c := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if c == '[' {
					break
				}
				s = append([]rune{c}, s...)
			}
			num := 0
			index := 1
			for {
				if len(stack) == 0 {
					break
				}
				c := stack[len(stack)-1]
				if c > '9' || c < '0' {
					break
				}
				stack = stack[:len(stack)-1]
				num = int(c-'0')*index + num
				index = index * 10
			}
			for i := 0; i < num; i++ {
				stack = append(stack, s...)
			}
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}
