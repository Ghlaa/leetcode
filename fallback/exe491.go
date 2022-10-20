package main

//var res [][]int
//func findSubsequences(nums []int) [][]int {
//    res=[][]int{}
//    dfs(nums,0,-101,[]int{})
//	return res
//}
//func dfs(nums []int,startindex int,curVal int,temp []int){
//	if len(temp)>=2{
//		tmp:=make([]int,len(temp))
//		copy(tmp,temp)
//		res=append(res,tmp)
//	}
//	hash:=make(map[int]bool,0)
//	for i:=startindex;i<len(nums);i++{
//		if nums[i]>=curVal{
//			if hash[nums[i]]==false{
//				hash[nums[i]]=true
//				temp=append(temp,nums[i])
//				linshi:=curVal
//				curVal=nums[i]
//				dfs(nums,i+1,curVal,temp)
//				temp=temp[:len(temp)-1]
//				curVal=linshi
//			}
//
//		}
//	}
//}