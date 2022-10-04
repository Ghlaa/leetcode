package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hashmap1 := map[int]int{}
	res := 0
	for _, a := range nums1 {
		for _, b := range nums2 {
			hashmap1[a+b]++
		}
	}
	for _, c := range nums1 {
		for _, d := range nums2 {
			if count, ok := hashmap1[0-c-d]; ok {
				res += count
			}
		}
	}
	return res
}
