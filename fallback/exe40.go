package fallback

//
//import (
//	"encoding/json"
//	"sort"
//)
//
//var res [][]int
//func combinationSum2(candidates []int, target int) [][]int {
//	res=[][]int{}
//	sort.Ints(candidates)
//	dfs(candidates,0,0,[]int{},target)
//	quchong(res)
//	return res
//}
//
//func dfs(candidates []int,index int,sum int,temp []int,target int){
//	//jainzhi
//
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
//		//if i>index&&candidates[i]==candidates[i-1]{
//		//	continue
//		//}
//		temp=append(temp,candidates[i])
//		sum+=candidates[i]
//		dfs(candidates,i+1,sum,temp,target)
//		temp=temp[0:len(temp)-1]
//		sum-=candidates[i]
//	}
//}
//
//func quchong(temp [][]int)[][]int{
//	 ans:=[][]int{}
//	/**
//	 *  这里是添加元素的操作.....
//	 *
//	 **/
//
//	// 开始去重
//	m := make(map[string]bool,len(temp))
//	for i := 0; i < len(temp); i++ {
//		marshal, _ := json.Marshal(temp[i])
//		// go 序列化后是 byte[]，需要转成 string
//		if m[string(marshal)]==false{
//			m[string(marshal)] = true
//		}
//
//	}
//
//	for key, _ := range m {
//		var item []int
//		json.Unmarshal([]byte(key), &item)
//		ans = append(ans, item)
//	}
//	return ans
//}
//
