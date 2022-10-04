package stringsExe

func romanToInt(s string) int {
	hashmap := make(map[byte]int)
	hashmap['I'] = 1
	hashmap['V'] = 5
	hashmap['X'] = 10
	hashmap['L'] = 50
	hashmap['C'] = 100
	hashmap['D'] = 500
	hashmap['M'] = 1000
	res := 0
	len := 0
	for i := 0; i < len; i++ {
		if i < (len-1) && (hashmap[s[i]] < hashmap[s[i+1]]) {
			res = res + hashmap[s[i+1]] - hashmap[s[i]]
			i++
		} else {
			res += hashmap[s[i]]
		}
	}
	return res
}
