/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	// 计算链表长度
	length := 0
	{
		_h := head
		for _h != nil {
			length++
			_h = _h.Next
		}
	}

	// 判断k是否大于链表长度
	if k%length == 0 {
		return head
	}
	if k > length {
		k = k % length
	}

	slow, fast := head, head

	for fast.Next != nil {
		fast = fast.Next
		k--
		if k < 0 {
			slow = slow.Next
		}
	}

	rtn := slow.Next
	slow.Next = nil
	fast.Next = head

	return rtn
}

// @lc code=end
