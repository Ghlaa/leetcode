package main

//var res [][]int
//func combinationSum3(k int, n int) [][]int {
//	res=[][]int{}
//	if k==0||n==0{
//		return res
//	}
//	fallback2(k,n,1,[]int{},0)
//	return res
//}
//
//func fallback2(k int,n int,start int,temp []int,sum int){
//	//fmt.Println(sum)
//	if len(temp)==k&&sum==n{
//		tmp:=make([]int,k)
//		copy(tmp,temp)
//		res=append(res,tmp)
//		return
//	}
//	//剪纸
//	if len(temp)>k||sum>n{
//		return
//	}
//	if len(temp)+9-start+1<k{
//		return
//	}
//
//	for i:=start;i<=9;i++{
//		sum+=i
//		temp=append(temp,i)
//		fallback2(k,n,i+1,temp,sum)
//		temp=temp[0:len(temp)-1]
//		sum-=i
//	}
//}

/*

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	sum := 0

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k && sum == n {
			ans = append(ans, append([]int{}, track...))
			return
		}
		if len(track) > k || sum > n {
			return
		}

		for i := start; i < 10; i++ {
			sum += i
			track = append(track, i)
			backtrack(i + 1)
			sum -= i
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return ans
}


*/