package tanxin

import "sort"

func maxSubArray(nums []int) int {
	lens:=len(nums)
	dp:=make([]int,lens)
	if lens==0{
		return 0
	}
	dp[0]=nums[0]
	for i:=1;i<lens;i++{
		if nums[i]>=nums[i]+dp[i-1]{
			dp[i]=nums[i]
		}else{
			dp[i]=nums[i]+dp[i-1]
		}
	}
	sort.Ints(dp)
	return dp[lens-1]
}
