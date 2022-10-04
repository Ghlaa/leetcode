package stringsExe

func longestPalindrome(s string) string {

	lens := len(s)
	if lens < 2 {
		return s //如果就一个字符直接返回
	}
	indexi := 0
	indexj := 0
	maxLen := 0
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			if isHuiWen(s[i:j+1]) == true && j-i+1 > maxLen {
				indexi = i
				indexj = j
			}
		}
	}
	return s[indexi : indexj+1]
}

func isHuiWen(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
