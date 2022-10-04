package stringsExe

func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return ""
	} else {
		res = strs[0]
	}

	for i := 1; i < len(strs); i++ {
		res = longestFix(res, strs[i])
	}
	return res
}

func longestFix(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	minn := 0
	if len1 > len2 {
		minn = len2
	} else {
		minn = len1
	}
	res := ""
	for i := 0; i < minn; i++ {
		//fmt.Println("是否")
		if str1[i] != str2[i] {
			//fmt.Println("测试:"+str1[:i]
			break
		}
		res = str1[:i+1]
	}
	return res
}

//标准答案中的分治法
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
