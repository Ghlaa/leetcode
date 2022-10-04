package stringsExe

import "strconv"

func addStrings1(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	var res, yushu, sum, jinwei int
	for i, j, k := len1-1, len2-1, 1; i >= 0 || j >= 0; i, j, k = i-1, j-1, k*10 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		sum = x + y + jinwei

		yushu = sum % 10
		jinwei = sum / 10
		if i <= 0 && j <= 0 {
			res = res + sum*k
		} else {
			res = res + yushu*k
		}

	}
	res1 := strconv.Itoa(res)
	return res1
}
