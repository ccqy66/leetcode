package two_point

import (
	"fmt"
	"github.com/ccqy66/leetcode/model"
	"testing"
)

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}

func TestThreeSumClosest(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 1))
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &model.ListNode{Val: 1,
		Next:&model.ListNode{Val: 2, Next:
			&model.ListNode{Val: 3, Next:
				&model.ListNode{Val: 4, Next:
					&model.ListNode{Val: 5, Next: nil}}}}}
	ans := removeNthFromEnd(head, 1)
	ans.Print()
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1,2}))
}