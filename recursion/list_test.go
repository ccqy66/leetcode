package recursion

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val:  5,
	}
	l2 := &ListNode{
		Val:  5,
	}
	result := addTwoNumbers(l1, l2)
	for item := result; item != nil;  item = item.Next{
		fmt.Println(item)
	}

}

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
			},
		},
	}
	l2 := &ListNode{
		Val:  5,
	}
	result := mergeTwoLists2(l1, l2)
	for item := result; item != nil;  item = item.Next{
		fmt.Println(item)
	}
}

func TestSwapPairs(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
			},
		},
	}
	result := swapPairs(l1)
	for item := result; item != nil;  item = item.Next{
		fmt.Println(item)
	}
}