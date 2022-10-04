package stringsExe

import (
	"fmt"
	"strconv"
)

var res []string
var hashmap map[int]string = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}

	huishu(digits, 0, "")
	return res
}

func huishu(digits string, index int, str string) {
	if index == len(digits) {
		res = append(res, str)
	} else {

		currentDig, _ := strconv.Atoi(string(digits[index]))

		fmt.Println(currentDig)
		fmt.Printf("%t \n", currentDig)
		currentstr := hashmap[currentDig] //表示当前的数字的字符串
		for i := 0; i < len(currentstr); i++ {
			huishu(digits, index+1, str+string(currentstr[i]))
		}
	}

}
