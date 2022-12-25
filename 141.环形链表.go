/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
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
func hasCycle(head *ListNode) bool {
	// 方法1: 使用hash
	// m := make(map[*ListNode]struct{})

	// for head != nil {
	// 	m[head] = struct{}{}
	// 	head = head.Next

	// 	if _, ok := m[head]; ok {
	// 		return true
	// 	}
	// }

	// return false

	// 方法2: 使用快慢指针
	slow, fast := head, head
	for step := 0; fast != nil; {
		fast = fast.Next
		step++

		if fast == slow {
			return true
		}

		if step%2 == 0 {
			slow = slow.Next
		}
	}
	return false
}

// @lc code=end
