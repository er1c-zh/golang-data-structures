package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lr := &ListNode{0, nil}
	lc := lr
	carry := 0
	for {
		lc.Val = l1.Val + l2.Val + carry
		carry = lc.Val / 10
		lc.Val = lc.Val % 10
		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil || l2 == nil {
			if l1 != nil {
				lc.Next = l1
				doCarry(lc, carry)
				return lr
			} else if l2 != nil {
				lc.Next = l2
				doCarry(lc, carry)
				return lr
			} else {
				if carry != 0 {
					lc.Next = &ListNode{carry, nil}
				}
				return lr
			}
		}
		lc.Next = &ListNode{0, nil}
		lc = lc.Next
	}

}

func doCarry(n *ListNode, carry int) {
	if carry != 0 {
		for {
			if n.Next == nil || carry == 0 {
				break
			}
			n.Next.Val += carry
			carry = n.Next.Val / 10
			n.Next.Val %= 10

			n = n.Next
		}

		if carry != 0 {
			n.Next = &ListNode{carry, nil}
		}
	}
}
