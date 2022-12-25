/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}

	slow, fast := head, head

	length := 1

	for fast.Next != nil {
		fast = fast.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	for i := 0; i < length-k-1; i++ {
		slow = slow.Next
	}

	next := slow.Next
	slow.Next = nil
	fast.Next = head

	return next
}

// @lc code=end
