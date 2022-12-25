/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := m[headA]; ok {
				return headA
			} else {
				m[headA] = struct{}{}
				headA = headA.Next
			}
		}

		if headB != nil {
			if _, ok := m[headB]; ok {
				return headB
			} else {
				m[headB] = struct{}{}
				headB = headB.Next
			}
		}

	}

	return nil
}

// @lc code=end
