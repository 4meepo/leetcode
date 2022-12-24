/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 */
package leetcode

import "container/list"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	rtn := [][]int{}

	deque := list.New()
	deque.PushBack(root)

	n := 1
	level := 1
	for {
		delta := 0

		arr := []int{}
		if level%2 == 0 {
			// <--
			for i := 0; i < n; i++ {
				node := deque.Remove(deque.Back()).(*TreeNode)
				arr = append(arr, node.Val)
				if node.Right != nil {
					deque.PushFront(node.Right)
					delta++
				}
				if node.Left != nil {
					deque.PushFront(node.Left)
					delta++
				}
			}
		} else {
			// -->
			for i := 0; i < n; i++ {
				node := deque.Remove(deque.Front()).(*TreeNode)
				arr = append(arr, node.Val)
				if node.Left != nil {
					deque.PushBack(node.Left)
					delta++
				}
				if node.Right != nil {
					deque.PushBack(node.Right)
					delta++
				}

			}

		}
		rtn = append(rtn, arr)
		if delta < 1 {
			break
		}
		n = delta
		level++
	}

	return rtn
}

// @lc code=end
