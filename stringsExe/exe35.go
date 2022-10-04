package stringsExe

func searchInsert(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	if left > right {
		return left
	} else {
		return right
	}
}
