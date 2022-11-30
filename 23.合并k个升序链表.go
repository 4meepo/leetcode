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

type nodePriorityQueue []*ListNode

func (hp nodePriorityQueue) Len() int {
	return len(hp)
}
func (hp nodePriorityQueue) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}
func (hp nodePriorityQueue) Less(i, j int) bool {
	return hp[i].Val < hp[j].Val
}
func (hp *nodePriorityQueue) Push(x interface{}) {
	*hp = append(*hp, x.(*ListNode))
}
func (hp *nodePriorityQueue) Pop() interface{} {
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

	queue := nodePriorityQueue{}
	for _, list := range lists {
		for list != nil {
			next := list.Next
			list.Next = nil
			queue.Push(list)
			list = next
		}
	}

	heap.Init(&queue)

	var head, cur *ListNode

	for queue.Len() > 0 {
		next := heap.Pop(&queue).(*ListNode)
		if head == nil {
			head = next
			cur = next
		} else {
			cur.Next = next
			cur = next
		}

	}

	return head
}

// @lc code=end
