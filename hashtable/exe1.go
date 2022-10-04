package main

func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	for index, m := range nums {
		if preindex, ok := hashmap[target-m]; ok {
			return []int{preindex, index}
		} else {
			hashmap[m] = index
		}
	}
	return []int{}
}
