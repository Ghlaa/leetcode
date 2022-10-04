package stringsExe

func calculate(s string) int {
	s = "+" + s
	tempNum := []int{}
	for i := 0; i < len(s)-1; i = i + 2 {
		switch s[i] {
		case '+':
			tempNum = append(tempNum, int(s[i+1]-'0'))
		case '-':
			tempNum = append(tempNum, int(s[i+1]-'0')*-1)
		case '*':
			tempint := tempNum[len(tempNum)-1] * int(s[i+1]-'0')
			tempNum[len(tempNum)-1] = tempint
		case '/':
			tempint := tempNum[len(tempNum)-1] / (int(s[i+1] - '0'))
			tempNum[len(tempNum)-1] = tempint
		}
	}
	sum := 0
	for _, x := range tempNum {
		sum += x
	}
	//fmt.Println(tempNum)
	return sum
}
