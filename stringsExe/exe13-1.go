package stringsExe

func romanToInt1(s string) int {
	hashmap := make(map[byte]int)
	hashmap['I'] = 1
	hashmap['V'] = 5
	hashmap['X'] = 10
	hashmap['L'] = 50
	hashmap['C'] = 100
	hashmap['D'] = 500
	hashmap['M'] = 1000
	res := 0
	for i := 1; i < len(s); i++ {
		if hashmap[s[i]] <= hashmap[s[i-1]] {
			res += hashmap[s[i-1]]
		} else {
			res -= hashmap[s[i-1]]
		}
	}
	res = res + hashmap[s[len(s)-1]]
	return res
}
