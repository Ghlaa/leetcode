package stringsExe

import "strconv"

const N = 4

var resaddress []string
var seg []int

func restoreIpAddresses(s string) []string {
	seg = make([]int, N)
	resaddress = []string{}
	dfs(s, 0, 0)
	return resaddress
}

func dfs(s string, segid int, start int) {

	if segid == N {
		if start == len(s) {
			ipstr := ""
			for i := 0; i < N; i++ {
				ipstr += strconv.Itoa(seg[i])
				if i != N-1 {
					ipstr += "."
				}
			}
			resaddress = append(resaddress, ipstr)
		}
		return
	}
	if start == len(s) {
		return
	}

	if s[start] == '0' {
		seg[segid] = 0
		dfs(s, segid+1, start+1)
	}

	//一般情况
	address := 0
	for i := start; i < len(s); i++ {
		address = address*10 + int(s[i]-'0')
		if address > 0 && address < 256 {
			seg[segid] = address
			dfs(s, segid+1, i+1)
		} else {
			break
		}
	}
}
