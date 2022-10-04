package stringsExe

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := ""
	lens := len(s)
	maxrows := 2 * (numRows - 1)
	for i := 0; i < numRows; i++ {
		for j, k := i, 1; j < lens; k++ {
			if i == 0 || i == numRows-1 {
				res = res + string(s[j])
				j = j + maxrows
			} else {
				if k%2 == 1 {
					res = res + string(s[j])
					j = j + maxrows - 2*i
				} else {
					res = res + string(s[j])
					j = j + 2*i
				}
			}
		}
	}
	return res
}
