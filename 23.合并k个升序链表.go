/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */
package main

import "container/heap"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type minHeap23 []*ListNode

func (hp minHeap23) Len() int {
	return len(hp)
}
func (hp minHeap23) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}
func (hp minHeap23) Less(i, j int) bool {
	return hp[i].Val < hp[j].Val
}
func (hp *minHeap23) Push(x interface{}) {
	*hp = append(*hp, x.(*ListNode))
}
func (hp *minHeap23) Pop() interface{} {
	old := *hp
	n := len(old)
	x := old[n-1]
	*hp = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	hp := make(minHeap23, 0, len(lists))

	// 用小顶堆维护k个链表的头节点的大小关系
	for _, node := range lists {
		if node != nil {
			heap.Push(&hp, node)
		}
	}

	var head, tail *ListNode
	for len(hp) > 0 {
		min := heap.Pop(&hp).(*ListNode)
		if head == nil {
			head = min
			tail = min
		} else {
			tail.Next = min
			tail = min
		}

		next := min.Next
		min.Next = nil
		if next != nil {
			heap.Push(&hp, next)
		}
	}

	return head
}

// @lc code=end
