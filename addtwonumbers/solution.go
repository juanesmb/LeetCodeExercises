package addtwonumbers

import "fmt"

func Init() {
	// [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	// [5,6,4]

	// [6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]

	// [9,9,9,9,9,9,9]
	// [9,9,9,9]

	// [8,9,9,9,0,0,0,1]

	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	result := addTwoNumbers(l1, l2)

	fmt.Println(result)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	index := result
	index = nil

	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		if index == nil {
			index = result
		} else {
			index.Next = &ListNode{}
			index = index.Next
		}

		index.Val, carry = calculateSum(l1, l2, carry)

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return result
}

func calculateSum(l1, l2 *ListNode, carry int) (int, int) {
	var d1 = 0
	var d2 = 0

	if l1 != nil {
		d1 = l1.Val
	}
	if l2 != nil {
		d2 = l2.Val
	}

	sum := d1 + d2 + carry
	newCarry := sum / 10
	if newCarry > 0 {
		return sum % 10, newCarry
	}

	return sum, newCarry
}
