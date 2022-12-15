/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode
	for list1 != nil && list2 != nil {
		if head == nil {
			head = tail
		}

		var target *ListNode
		if list1.Val <= list2.Val {
			target = list1
			list1 = list1.Next
		} else {
			target = list2
			list2 = list2.Next
		}

		if tail == nil {
			tail = target
		} else {
			tail.Next = target
			tail = tail.Next
		}
	}

	// list1 remain
	if list1 != nil {
		if tail == nil {
			// list2 为空
			return list1
		}
		tail.Next = list1
	}
	// list2 remain
	if list2 != nil {
		if tail == nil {
			// list1 为空的情况
			return list2
		}
		tail.Next = list2
	}

	if head == nil {
		return tail
	}

	return head
}

// @lc code=end
