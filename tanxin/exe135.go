package tanxin

import (
	"fmt"
	"math"
)

func candy(ratings []int) int {
	dp := make([]int, len(ratings))
	dp[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			dp[i] = int(math.Max(float64(dp[i+1]+1), float64(dp[i])))
		}
	}
	sum := 0
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
		sum += dp[i]
	}
	return sum
}
