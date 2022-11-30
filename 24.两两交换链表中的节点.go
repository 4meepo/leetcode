/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// length == 0
	if head == nil {
		return nil
	}
	// length == 1
	if head.Next == nil {
		return head
	}

	var newHead, s, d, lastS *ListNode
	for i := 1; head != nil; i++ {
		cur := head
		if i%2 == 0 {
			// 偶数
			d = cur
		} else {
			// 奇数
			s = cur
		}

		if newHead == nil {
			newHead = d
		}

		// swap
		head = cur.Next
		if i%2 == 0 {
			d.Next = s
			s.Next = head

			if i > 2 {
				lastS.Next = d
			}

			lastS = s
			// lastD = d
		}

	}

	return newHead
}

// @lc code=end
