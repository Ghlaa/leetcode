package main
//var res [][]int
//func combine(n int, k int) [][]int {
//
//	temp:=[]int{}
//	fallback(n,k,1,temp)
//	return res
//}
//
//func fallback(n int, k int,start int,temp []int){
//	if len(temp)==k{
//		res=append(res,temp)
//		return
//	}
//	for i:=start;i<n;i++{
//		temp=append(temp,i)
//		fallback(n,k,start+1,temp)
//		temp=temp[0:len(temp)-1]
//	}
//}