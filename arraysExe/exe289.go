package main

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	if m == 0 {
		return
	}
	//记录八个方向
	left := [8]int{-1, 0, 1, 1, 1, 0, -1, -1}
	right := [8]int{-1, -1, -1, 0, 1, 1, 1, 0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := 0
			for k := 0; k < 8; k++ {
				x := i + left[k]
				y := j + right[k]
				if x >= 0 && x < n && y >= 0 && y < m {
					cnt += board[x][y] & 1 //注意我们判断周围细胞状态的时候，只需要考虑二进制数的最后一位，1其实就代表01，位运算
				}
				if board[i][j]&1 > 0 { //证明此时是活细胞
					if cnt == 2 || cnt == 3 { //如果周围还有两个或者三个活细胞的话，下一个状态还是活细胞
						board[i][j] = 0b11 //高位的1代表细胞的下一个状态。低位的1代表当前状态，其他细胞的状态变化需要参考的是低位的当前状态
					}
				} else {
					//
					if cnt == 3 {
						board[i][j] = 0b10 //高位的1代表细胞的下一个状态。低位的0代表当前状态，其他细胞的状态变化需要参考的是低位的当前状态
					}
				}
			}
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				board[i][j] = board[i][j] >> 1 //将一个二进制数进行右移，相当于舍弃了低位，改成了高位，即细胞的下一个状态
			}
		}
	}
}
