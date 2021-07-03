package recursion

import (
	"fmt"
	"strings"
)

func Permutation(str string) {
	strArray := make([]uint8, 0, len(str))
	for i:=0; i< len(str); i++ {
		strArray = append(strArray, str[i])
	}
	permutation(strArray, 0, len(str))
}
func permutation(str []uint8, start, end int) {
	if start >= end {
		println(string(str))
		return
	}
	for i := start; i < end; i++ {
		str[start], str[i] = str[i], str[start]
		permutation(str, start+1, end)
		str[start], str[i] = str[i], str[start]
	}
}
var ans int
func longestSubstring(s string, k int) int {
	ans = 0
	longestString(s, k)
	return ans
}

func longestString(s string, k int) int{
	if len(s) < k {
		return 0
	}
	m := make(map[uint8]int)
	for i:=0; i< len(s);i++ {
		m[s[i]]++
	}
	for c, v := range m{
		if v < k {
			var tmp int
			for _,item :=  range strings.Split(s, string(c)){
				tmp = Max(tmp, longestString(item, k))
			}
			ans = Max(ans, tmp)
			return tmp
		}
	}
	ans = Max(ans, len(s))
	return len(s)
}

func kthGrammar(n int, k int) int {
	return int(grammar(n)[k-1] - '0')
}

func grammar(n int) string {
	if n == 1 {
		return "0"
	}
	if n == 2 {
		return "01"
	}

	tmp := grammar(n-1)
	ans := tmp
	for i:=len(tmp)/2; i< len(tmp); i++{
		if tmp[i] == '0' {
			ans = fmt.Sprintf("%s%s", ans, "01")
		}else {
			ans = fmt.Sprintf("%s%s", ans, "10")
		}
	}
	return ans
}
