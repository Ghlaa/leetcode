package stringsExe

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	hashmap := make(map[byte]byte)
	hashmap['('] = ')'
	hashmap['['] = ']'
	hashmap['{'] = '}'
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if hashmap[s[i]] == 0 {
			if len(stack) == 0 || hashmap[stack[len(stack)-1]] != hashmap[s[i]] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, hashmap[s[i]])
		}
	}
	if len(stack) != 0 {
		return false
	} else {
		return true
	}
}
