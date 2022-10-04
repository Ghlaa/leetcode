package stringsExe

func maxArea1(height []int) int {
	maxSize := 0
	heightSize := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[j] >= height[i] {
				heightSize = height[i]
			} else {
				heightSize = height[j]
			}
			if (j-i)*heightSize > maxSize {
				maxSize = (j - i) * heightSize
			}
		}
	}
	return maxSize
}

func testF(a, b int) int {
	return a + b
}
