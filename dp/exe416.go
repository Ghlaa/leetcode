package dp

import "sort"

func canPartition(nums []int) bool {
	sum:=0
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
	}
	if sum%2==1{
		return false
	}
	sum=sum/2
	dp := make([]int, len(nums)+1)


	for i:=0;i<len(nums);i++{
		for j:=sum;j>=nums[i];j--{
			dp[j]=max(dp[j],dp[j-nums[i]]+nums[i])
		}
	}
	if  dp[sum]==sum{
		return true
	}else {
		return false
	}
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}