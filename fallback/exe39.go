package fallback

//var res [][]int
//func combinationSum(candidates []int, target int) [][]int {
//	res=[][]int{}
//	dfs(candidates,0,0,[]int{},target)
//	return res
//}
//
//func dfs(candidates []int,index int,sum int,temp []int,target int){
//	//jainzhi
//	if sum>target{
//		return
//	}
//
//	//保存正确结果
//	if sum==target{
//		tmp:=make([]int,len(temp))
//		copy(tmp,temp)
//		res=append(res,tmp)
//	}
//	for i:=index;i<len(candidates);i++{
//		temp=append(temp,candidates[i])
//		sum+=candidates[i]
//		dfs(candidates,i,sum,temp,target)
//		temp=temp[0:len(temp)-1]
//		sum-=candidates[i]
//	}
//}
