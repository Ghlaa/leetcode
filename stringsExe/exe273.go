package stringsExe

import "strings"

var (
	singles   = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens      = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands = []string{"", "Thousand", "Million", "Billion"}
)
var sb = &strings.Builder{}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	for i, unit := 3, int(1e9); i >= 0; i-- {
		if curNum := num / unit; curNum > 0 {
			num -= curNum * unit
			recursion(curNum)
			sb.WriteString(thousands[i])
			sb.WriteByte(' ')
		}
		unit /= 1000
	}
	return strings.TrimSpace(sb.String())
}

func recursion(num int) {
	switch {
	case num == 0:
	case num < 10:
		sb.WriteString(singles[num])
		sb.WriteByte(' ')
	case num < 20:
		sb.WriteString(teens[num-10])
		sb.WriteByte(' ')
	case num < 100:
		sb.WriteString(tens[num/10])
		sb.WriteByte(' ')
		recursion(num % 10)
	default:
		sb.WriteString(singles[num/100])
		sb.WriteString(" Hundred ")
		recursion(num % 100)
	}
}
