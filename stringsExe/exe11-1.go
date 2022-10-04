package stringsExe

func maxArea(height []int) int {
	j := len(height) - 1
	i := 0
	maxSize := 0

	for i <= j {
		if height[i] <= height[j] {
			if height[i]*(j-i) > maxSize {
				maxSize = height[i] * (j - i)
			}
			i++
		} else {
			if height[j]*(j-i) > maxSize {
				maxSize = height[j] * (j - i)
			}
			j--
		}
	}
	return maxSize
}
