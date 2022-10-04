package main

func isAnagram(s string, t string) bool {
	len1 := len(s)
	len2 := len(t)
	if len1 != len2 {
		return false
	}
	hashtable1 := map[byte]int{}
	for i := 0; i < len1; i++ {
		hashtable1[s[i]]++
	}
	hashtable2 := map[byte]int{}
	for i := 0; i < len2; i++ {
		hashtable2[t[i]]++
	}
	for i := 0; i < len2; i++ {
		if hashtable1[t[i]] != hashtable2[t[i]] {
			return false
		}
	}
	return true
}
