package tanxin
func maxProfit(prices []int) int {
	lens:=len(prices)
	if lens==0||lens==1{
		return 0
	}
	dp:=make([]int,lens)
	dp[0]=0
	for i:=1;i<lens;i++{
		if prices[i]-prices[i-1]>0{
			dp[i]=prices[i]-prices[i-1]+dp[i-1]
		}else{
			dp[i]=dp[i-1]
		}
	}

	return dp[lens-1]
}

