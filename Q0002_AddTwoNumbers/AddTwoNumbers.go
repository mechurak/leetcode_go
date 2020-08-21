package q0002

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	up := 0
	var retList []*ListNode
	for l1 != nil || l2 != nil {
		left := 0
		right := 0
		sum := up
		if l1 != nil {
			left = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			right = l2.Val
			l2 = l2.Next
		}
		sum = sum + left + right
		if sum >= 10 {
			up = 1
			sum = sum % 10
		} else {
			up = 0
		}
		retList = append(retList, &ListNode{Val: sum, Next: nil})
	}
	if up > 0 {
		retList = append(retList, &ListNode{Val: up, Next: nil})
	}
	for i := 0; i < len(retList)-1; i++ {
		retList[i].Next = retList[i+1]
	}
	return retList[0]
}
