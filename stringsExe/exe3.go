package stringsExe

//func lengthOfLongestSubstring(s string) int {
//	maxLen:=0
//	lens:=len(s)
//	for i:=0;i<lens;i++{
//		if maxLen>=lens-i{
//			return maxLen
//		}
//		for j:=i;j<lens;j++{
//			if judge(s[i:j+1])==true&&j-i+1>maxLen{
//				maxLen=j-i+1
//			}
//		}
//	}
//	return maxLen
//}

//func lengthOfLongestSubstring(s string) int {
//
//	lens:=len(s)
//	if lens==0 {
//		return 0
//	}
//	maxLen:=1//至少为1
//	for L:=2;L<=lens;L++{//L为长度
//		for i:=0;i+L<=lens;i++{
//			if judge(s[i:i+L])==true{
//				maxLen=L
//				continue
//			}
//		}
//	}
//	return maxLen
//}
//func judge(s string) bool{
//	hashmap:=make(map[byte]int)
//	for i:=0;i<len(s);i++{
//		if hashmap[s[i]]==0{
//			hashmap[s[i]]++
//		}else{
//			return false
//		}
//	}
//	return true
//}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
