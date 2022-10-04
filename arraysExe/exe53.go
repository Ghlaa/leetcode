package main

import "fmt"

func maxSubArray(nums []int) int {
	lens := len(nums)
	dp := make([]int, lens)
	dp[0] = nums[0]
	for i := 1; i < lens; i++ {
		if nums[i]+dp[i-1] > nums[i] {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
	}
	res := nums[0]
	for i := 0; i < lens; i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	fmt.Println(dp)
	return res
}
