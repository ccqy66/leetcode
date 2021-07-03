package recursion

import (
	"fmt"
	"testing"
)

func TestSameDifference(t *testing.T) {
	result := sameDifference(2,1)
	fmt.Println(result)
}

func TestBack2Zero(t *testing.T) {
	fmt.Println(back2Zero([]int{4,2,3,0,3,1,2}, 5))
}

func TestDivideBoard(t *testing.T) {
	fmt.Println(divideBoard(1,2,3))
}

func TestSearchLeft(t *testing.T) {
	println(searchLeft([]int{1,3,3,3,4,5}, 2))
}

func TestSearchRight(t *testing.T) {
	println(searchRight([]int{1,2,2,4,5,6}, 2))
}