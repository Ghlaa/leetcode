package stringsExe

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] == s[j] {
			i++
			j--
		} else {
			if (isHuiwen(s[i+1:j+1])) == true || (isHuiwen(s[i:j])) == true {
				return true
			} else {
				return false
			}
		}
	}
	return true
}
func isHuiwen(a string) bool {
	lens := len(a)
	for i := 0; i < lens/2; i++ {
		if a[i] != a[lens-1-i] {
			return false
		}
	}
	return true
}
