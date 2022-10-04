package main

func canConstruct(ransomNote string, magazine string) bool {
	hashmap := map[byte]int{}
	for i := 0; i < len(ransomNote); i++ {
		hashmap[ransomNote[i]]++
	}
	for i := 0; i < len(magazine); i++ {
		hashmap[magazine[i]]--
		if hashmap[magazine[i]] < 0 {
			return false
		}
	}
	return true
}
