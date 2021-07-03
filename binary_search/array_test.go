package binary_search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{3,1}, 1))
	//fmt.Println(search([]int{5,6,7,0,1,2,3,4}, 1))
	//fmt.Println(search([]int{5,6,7,0,1,2,3,4}, 2))
	//fmt.Println(search([]int{5,6,7,0,1,2,3,4}, 3))
	//fmt.Println(search([]int{5,6,7,0,1,2,3,4}, 4))
	//fmt.Println(search([]int{5,6,7,0,1,2,3,4}, 5))
	//fmt.Println(search([]int{5,6,7,0,1,2,3,4}, 6))
	//fmt.Println(search([]int{5,6,7,0,1,2,3,4}, 7))

}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}

func TestFirstBadVersion(t *testing.T) {
	fmt.Println(firstBadVersion(5))
}

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{2,2}, 3))
}

func TestSearchMatrix(t *testing.T){
	s := [][]int{{1,1}}
	fmt.Println(searchMatrix(s, 2))
}

func TestSearchV2(t *testing.T) {
	fmt.Println(searchV2([]int{1,0,1,1,1},0))
}

func TestFindMin(t *testing.T) {
	fmt.Println(findMin([]int{2,1}))
}