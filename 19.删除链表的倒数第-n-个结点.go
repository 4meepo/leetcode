/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}

	var slow, fast = head, head
	for fast.Next != nil {
		fast = fast.Next
		n--
		if n < 0 {
			slow = slow.Next
		}
	}
	// 如果链表长度刚好 == n
	if n > 0 {
		return slow.Next
	}
	slow.Next = slow.Next.Next

	return head
}

// @lc code=end
