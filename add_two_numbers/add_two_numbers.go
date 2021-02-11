package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	cur := head

	for l1 != nil || l2 != nil {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry

		carry = 0
		if sum >= 10 {
			carry = 1
		}

		resval := sum
		if resval >= 10 {
			resval -= 10
		}

		cur.Next = &ListNode{Val: resval}
		cur = cur.Next
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return head.Next
}
