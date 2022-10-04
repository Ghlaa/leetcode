package main

func rotate(matrix [][]int) {
	lens := len(matrix)
	temp := 0
	//先对矩阵进行上下转置
	for i := 0; i < lens/2; i++ {
		for j := 0; j < lens; j++ {
			temp = matrix[i][j]
			matrix[i][j] = matrix[lens-i-1][j]
			matrix[lens-i-1][j] = temp
		}
	}
	//对矩阵进行左对角线的转置
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ { //只需要对右上角的元素进行反转即可
			temp = matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}
