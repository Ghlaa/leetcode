package stackAndQueue

func removeDuplicates(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else if s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else if s[i] != stack[len(stack)-1] {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
