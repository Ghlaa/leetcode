package stringsExe

import "sort"

//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/group-anagrams
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//func groupAnagrams(strs []string) [][]string {
//	res:=[][] string{}
//	hashmap:=make(map[[26]int][]string)
//	for _,m:=range strs{//m是一个字符串
//		tempmap:=[26]int{}
//		//tempstring:=""
//		for i:=0;i<len(m);i++{
//			tempmap[m[i]-'a']++
//		}
//		//hashmap[tempmap]=m
//		hashmap[tempmap]=append(hashmap[tempmap],m)
//	}
//	for _,val:=range hashmap{
//		res=append(res,val)
//	}
//	return res
//}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	hashmap := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(x, y int) bool { return s[x] > s[y] })
		tempstr := string(s)
		hashmap[tempstr] = append(hashmap[tempstr], str)
	}
	for _, val := range hashmap {
		res = append(res, val)
	}
	return res
}
