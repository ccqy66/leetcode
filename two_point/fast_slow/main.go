package fast_slow

import "github.com/ccqy66/leetcode/model"

// 判断链表是否有环
func hasCycle(node *model.ListNode) bool {
	fast, slow := node, node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 成环链表的入口
func cycleEnter(node *model.ListNode) *model.ListNode {
	fast, slow := node, node
	var enter *model.ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			enter = fast
			fast = node
		}
	}
	if enter == nil {
		return nil
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func getLastKNode(node *model.ListNode, k int) *model.ListNode {
	fast, slow := node, node
	i := 0
	for ; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	// 说明链表的长度<k
	if i < k {
		return nil
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func getMiddleNode(node *model.ListNode) *model.ListNode {
	fast, slow := node, node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
