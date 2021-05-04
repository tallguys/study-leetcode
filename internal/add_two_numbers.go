package internal

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode-cn.com/problems/add-two-numbers/
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	cur := ret
	nextL1 := l1
	nextL2 := l2
	carry := 0
	for {
		if nextL1 != nil && nextL2 != nil {
			cur.Val = nextL1.Val + nextL2.Val + carry
			nextL1 = nextL1.Next
			nextL2 = nextL2.Next
		} else if nextL1 != nil && nextL2 == nil {
			cur.Val = nextL1.Val + carry
			nextL1 = nextL1.Next
			nextL2 = nil
		} else if nextL1 == nil && nextL2 != nil {
			cur.Val = nextL2.Val + carry
			nextL1 = nil
			nextL2 = nextL2.Next
		} else {
			cur.Val = carry
			nextL1 = nil
			nextL2 = nil
		}

		if cur.Val >= 10 {
			cur.Val = cur.Val - 10
			carry = 1
		} else {
			carry = 0
		}

		if nextL1 == nil && nextL2 == nil && carry == 0 {
			break
		}

		cur.Next = new(ListNode)
		cur = cur.Next
	}

	return ret
}
