package dp

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func TestMinPathSum(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1,3,1},{1,5,1},{4,2,1}}))
}
func TestNumDecodings(t *testing.T) {
	fmt.Println(numDecodings("1123"))
}
func TestIsInterleave(t *testing.T) {
	fmt.Println(isInterleave("", "", ""))
}

func TestMinimumTotal(t *testing.T) {
	fmt.Println(minimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}}))
}

func TestMaxLengthPath(t *testing.T) {
}

func TestFindChange(t *testing.T) {
	fmt.Println(findChange(27))
}