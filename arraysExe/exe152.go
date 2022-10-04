package main

import (
	"sort"
)

func maxProduct(nums []int) int {
	cnt := 0
	cnt_0 := 0
	res := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			cnt++
		}
		if nums[i] == 0 {
			cnt_0++
		}
		res *= nums[i]
	}
	if cnt%2 == 0 && cnt_0 == 0 {
		return res
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]*nums[i] > nums[i] {
			dp[i] = dp[i-1] * nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	sort.Ints(dp)
	res = dp[len(dp)-1]
	return res
}
