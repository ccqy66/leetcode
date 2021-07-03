package stack

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	fmt.Println(removeDuplicateLetters("bcabc"))
}

func TestNextPermutation(t *testing.T) {
	arr := []int{1 ,2, 3, 5, 4, 2}
	nextPermutation(arr)
	fmt.Println(arr)
}