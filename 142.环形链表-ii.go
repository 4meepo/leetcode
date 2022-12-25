/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
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
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	// 用hash记录, 第二次访问到同一个节点时, 说明就是这里
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		} else {
			m[head] = struct{}{}
		}

		head = head.Next
	}

	return nil
}

// @lc code=end
