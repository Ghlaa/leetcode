package stringsExe

import "strconv"

func addBinary(a string, b string) string {
	len1, len2 := len(a), len(b)
	var yushu, sum, jinwei int
	resstr := ""
	for i, j := len1-1, len2-1; i >= 0 || j >= 0 || jinwei != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(a[i] - '0')
		}
		if j >= 0 {
			y = int(b[j] - '0')
		}

		sum = x + y + jinwei

		yushu = sum % 2
		jinwei = sum / 2
		//res=res+yushu*k
		resstr = strconv.Itoa(yushu) + resstr

	}
	return resstr
}
