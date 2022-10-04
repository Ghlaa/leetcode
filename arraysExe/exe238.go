package main

func productExceptSelf(nums []int) []int {

	lens := len(nums)
	left := make([]int, lens)
	left[0] = 1
	for i := 1; i < lens; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right := make([]int, lens)
	right[lens-1] = 1
	for i := lens - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	res := []int{}
	for i := 0; i < lens; i++ {
		res[i] = left[i] * right[i]
	}
	return res
}
