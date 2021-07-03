package windows

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("bbbbbb"))
}

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("a", "b"))
}

func TestMaxSlidingWindow(t *testing.T) {
	//fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(maxSlidingWindow([]int{7,2,4}, 2))
}

func TestCharacterReplacement(t *testing.T) {
	fmt.Println(characterReplacement("ABAA",0))
}

func TestCheckInclusion(t *testing.T) {
	fmt.Println(checkInclusion("ab", "abdacbooo"))
}

func TestLongestOnes(t *testing.T) {
	//							  0 1 2 3 4 5 6 7 8 9 10 11,12 13 14 15 16 17 18
	fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1, 1,  0, 0, 0, 1, 1, 1, 1}, 3))
}

func TestShortestSeq(t *testing.T) {
	fmt.Println(shortestSeq([]int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}, []int{1,5,9}))
}

func TestMinOperations(t *testing.T) {
	fmt.Println(minOperations([]int{8828,9581,49,9818,9974,9869,9991,10000,10000,10000,9999,9993,9904,8819,1231,6309},134365))
}