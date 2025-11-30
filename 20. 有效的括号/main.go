package main

func main() {
	isValid("()[]{}")
}
func isValid(s string) bool {
	stack := []byte{}
	for _, b := range s {
		stack = append(stack, byte(b))
		if len(stack) == 1 {
			continue
		} else {
			if equalRune(stack[len(stack)-2], stack[len(stack)-1]) {
				stack = stack[:len(stack)-2]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
func equalRune(a, b byte) bool {
	if a == '[' && b == ']' {
		return true
	}
	if a == '(' && b == ')' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	return false
}
