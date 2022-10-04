package stringsExe

// " aa bb"
func lengthOfLastWord(s string) int {
	lens := len(s) - 1
	res := 0
	for j := lens; j >= 0; j-- {
		if s[j] == ' ' {
			lens--
		} else {
			break
		}
	}
	for j := lens; j >= 0; j-- {
		if s[j] != ' ' {
			res++
		} else {
			break
		}
	}
	return res
}
