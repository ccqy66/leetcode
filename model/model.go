package model

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	head := l
	for head != nil {
		if head.Next == nil {
			fmt.Println(fmt.Sprintf("%d", head.Val))
		}else {
			fmt.Print(fmt.Sprintf("%d->", head.Val))
		}
		head = head.Next
	}
}
