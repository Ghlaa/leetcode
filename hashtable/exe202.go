package main

import "strconv"

func isHappy(n int) bool {
	var temp int
	hashmap := map[int]int{}
	for n != 1 {
		if hashmap[n] == 0 {
			hashmap[n] = 1
		} else {
			return false
		}
		m := strconv.Itoa(n)
		temp = 0
		for i := 0; i < len(m); i++ {
			temp += square(int(m[i] - '0'))
		}
		n, _ = strconv.Atoi(m)
	}
	return true
}

func square(n int) int {
	return n * n
}
