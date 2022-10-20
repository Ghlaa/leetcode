package main

//import "sort"
//
//var res [][]int
//func subsets(nums []int) [][]int {
//	sort.Ints(nums)
//	res=[][]int{}
//	dfs(nums,0,[]int{})
//	return res
//}
//func dfs(nums []int,start int,temp []int){
//	tmp:=make([]int,len(temp))
//	copy(tmp,temp)
//	res=append(res,tmp)
//	for i:=start;i<len(nums);i++{
//		if i>start&&nums[i]==nums[i-1]{
//			continue
//		}
//		temp=append(temp,nums[i])
//		dfs(nums,i+1,temp)
//		temp=temp[:len(temp)-1]
//	}
//}