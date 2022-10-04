package main

import "fmt"

//主goroutine
func main() {
	matrix := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(matrix, 7))
}
