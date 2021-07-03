package recursion

import (
	"fmt"
	"testing"
)

func TestPermutation2(t *testing.T) {
	r := permute([]int{1,2,3,4})
	fmt.Print(r)
}

func t1(s []int) {
	//s = s[:len(s)-1]
	s = append(s, 6)
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}
/**
1 1 2 5 6 7 10
|
1 2 5 6 7
|
6 7 10

 */
func TestCombinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
}