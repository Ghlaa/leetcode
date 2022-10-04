package main

func intersection(nums1 []int, nums2 []int) []int {
	hashmap := map[int]int{}
	res := []int{}
	for i := 0; i < len(nums1); i++ {
		hashmap[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		if hashmap[nums2[i]] == 1 {
			res = append(res, nums2[i])
			hashmap[nums2[i]] = 0
		}
	}
	return res
}
