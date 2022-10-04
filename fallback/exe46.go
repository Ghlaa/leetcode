package fallback

//var res [][]int
//var hash map[int]bool
//func permute(nums []int) [][]int {
//	res=[][]int{}
//	hash=make(map[int]bool,0)
//	dfs(nums,0,[]int{})
//	return res
//}
//func dfs(nums []int,count int,temp []int){
//	if count==len(nums){
//		tmp:=make([]int,len(temp))
//		copy(tmp,temp)
//		res=append(res,tmp)
//		return
//	}
//	for i:=0;i<len(nums);i++{
//		if hash[nums[i]]==false{
//			hash[nums[i]]=true
//			temp=append(temp,nums[i])
//			count++
//			dfs(nums,count,temp)
//			count--
//			temp=temp[0:len(temp)-1]
//			hash[nums[i]]=false
//		}
//
//	}
//}
