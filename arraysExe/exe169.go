package main

func majorityElement(nums []int) int {
	lens := len(nums)
	hashmap := make(map[int]int)
	for i := 0; i < lens; i++ {
		hashmap[nums[i]]++
		if hashmap[nums[i]] > lens/2 {
			return nums[i]
		}
	}
	for i := 0; i < len(hashmap); i++ {
		if hashmap[i] > lens/2 {
			return hashmap[i]
		}
	}
	return 0
}
