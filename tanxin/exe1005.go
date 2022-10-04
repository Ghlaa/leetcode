package tanxin

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sum := 0
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(math.Abs(float64(nums[j])))
	})
	for i := 0; i < len(nums) && k != 0; i++ {
		if nums[i] < 0 {
			nums[i] = nums[i] * -1
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] = -1 * nums[len(nums)-1]
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
