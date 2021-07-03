package recursion

import (
	"fmt"
	"testing"
)

func TestPermutation(t *testing.T) {
	Permutation("abc")
}

func TestLongestSubstring(t *testing.T) {
	fmt.Println(longestSubstring("aaabb", 2))
	//ss := strings.Split("aaab", string('b'))
	//fmt.Println(ss)
}

func TestKthGrammar(t *testing.T) {
	//fmt.Println(kthGrammar(1,1))
	//fmt.Println(kthGrammar(2,1))
	//fmt.Println(kthGrammar(2,2))
	//fmt.Println(kthGrammar(4,5))
	for i:=1; i< 10; i ++{
		fmt.Println(len(grammar(i)))
	}
	fmt.Println(grammar(1))
	fmt.Println(grammar(2))
	fmt.Println(grammar(3))
	fmt.Println(grammar(4))
	fmt.Println(grammar(5))
	fmt.Println(grammar(6))
	fmt.Println(grammar(7))
	fmt.Println(grammar(8))
}