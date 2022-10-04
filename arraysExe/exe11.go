package main

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	maxVal := 0
	temp := 0
	for i <= j {
		if height[i] >= height[j] {
			temp = (j - i) * height[j]
			if temp > maxVal {
				maxVal = temp
			}
			j--
		} else {
			temp = (j - i) * height[i]
			if temp > maxVal {
				maxVal = temp
			}
			i++
		}

	}
	return maxVal
}
