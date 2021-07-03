package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	result := &ListNode{}
	carry := 0
	var newList *ListNode
	if l1 == nil {
		result.Val = l2.Val
		newList = addTwoNumbers(nil, l2.Next)
	} else if l2 == nil {
		result.Val = l1.Val
		newList = addTwoNumbers(l1.Next, nil)
	} else {
		item := l1.Val + l2.Val
		carry = item / 10
		result.Val = item % 10
		newList = addTwoNumbers(l1.Next, l2.Next)
	}
	result.Next = newList
	var current *ListNode
	var last = result
	for current = newList; current != nil; current = current.Next {
		if carry > 0 {
			p := current.Val + carry
			current.Val = p % 10
			carry = p / 10
		}
		if current.Next == nil {
			last = current
		}
	}
	if carry > 0 {
		last.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return result
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode)*ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l2, l1.Next)
		return l1
	}else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var bigger *ListNode
	var smaller *ListNode

	if l1.Val >= l2.Val {
		bigger = l1
		smaller = l2
	}else {
		bigger = l2
		smaller = l1
	}
	if smaller.Next != nil && smaller.Next.Val <= bigger.Val {
		smaller.Next = mergeTwoLists2(bigger, smaller.Next)
	}else {
		tmp := smaller.Next
		smaller.Next = bigger
		if tmp != nil {
			bigger.Next = mergeTwoLists2(tmp, bigger.Next)
		}
	}

	return smaller
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := head.Next.Next
	ans := head.Next
	head.Next.Next = head
	head.Next = swapPairs(nextHead)
	return ans
}