package stackAndQueue

func isValid(s string) bool {

	hashmap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || hashmap[stack[len(stack)-1]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
