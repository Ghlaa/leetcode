package stackAndQueue

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	temp := 0
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			temp = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, temp)
		} else if tokens[i] == "-" {
			temp = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, temp)
		} else if tokens[i] == "*" {
			temp = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, temp)
		} else if tokens[i] == "/" {
			temp = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, temp)
		} else {
			interger, _ := strconv.Atoi(tokens[i])
			stack = append(stack, interger)
		}
	}
	return stack[0]
}
