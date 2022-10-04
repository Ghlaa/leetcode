package main

func spiralOrder(matrix [][]int) []int {
	resArr := []int{}
	m := len(matrix)
	n := len(matrix[0])
	if m == 0 || n == 0 {
		return resArr
	}
	xint := [4]int{0, 1, 0, -1}
	yint := [4]int{1, 0, -1, 0}

	count := m * n
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	x, y := 0, 0
	temp := 0
	xindex, yindex := 0, 0
	for i := 0; i < count; i++ {
		temp = matrix[x][y]
		resArr = append(resArr, temp)
		visited[x][y] = true
		newX, newY := x+xint[xindex], y+yint[yindex]
		if newX < 0 || newY < 0 || newX >= m || newY >= n || visited[newX][newY] == true {
			xindex = (xindex + 1) % 4
			yindex = (yindex + 1) % 4
		}
		x = x + xint[xindex]
		y = y + yint[yindex]
	}
	return resArr
}
