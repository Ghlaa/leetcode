package main

func plusOne(digits []int) []int {
	res := []int{}
	add := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if add == 1 && digits[i] == 9 {
			res = append(res, 0)
			add = 1
		} else {
			res = append(res, digits[i]+add)
			add = 0
		}
	}
	if add == 1 {
		res = append(res, 1)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
