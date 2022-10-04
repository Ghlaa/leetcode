package stringsExe

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || len(haystack) == 0 {
		return 0
	}

	m, n := len(haystack), len(needle)
	tem := 0
	flag := false
	for i := 0; i <= m-n; i++ {
		tem = i
		for j := 0; j < n; j++ {
			if haystack[tem] != needle[j] {
				flag = false
				break
			} else { //如果匹配的字符串前面几个相同的话
				tem++
				flag = true
			}
		}
		if flag == true {
			return i
		}

	}
	return -1
}
