/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	var overflow int
	var current *ListNode
	for l1 != nil || l2 != nil {
		if current == nil {
			current = head
		} else {
			current.Next = new(ListNode)
			current = current.Next
		}
		sum := 0
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		sum += overflow
		overflow = 0
		current.Val = sum % 10
		if sum >= 10 {
			overflow = 1
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if overflow == 1 {
		current.Next = &ListNode{Val: 1}
	}

	return head
}

// @lc code=end
