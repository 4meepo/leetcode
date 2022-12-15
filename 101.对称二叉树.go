/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	deque := list.New()

	deque.PushBack(root.Left)
	deque.PushBack(root.Right)

	for deque.Len() != 0 {
		if deque.Len()%2 != 0 {
			return false
		}

		f := deque.Front()
		b := deque.Back()
		for i := 0; i < deque.Len()/2; i++ {
			fNode := f.Value.(*TreeNode)
			bNode := b.Value.(*TreeNode)
			if (fNode == nil && bNode != nil) || (fNode != nil && bNode == nil) {
				// 其1为nil
				return false
			}
			if fNode == bNode {
				// 同为nil
			} else if fNode.Val != bNode.Val {
				return false
			}
			f = f.Next()
			b = b.Prev()
		}

		l := deque.Len()
		for i := 0; i < l; i++ {
			e := deque.Remove(deque.Front()).(*TreeNode)
			if e != nil {
				deque.PushBack(e.Left)
			}
			if e != nil {
				deque.PushBack(e.Right)
			}
		}

	}

	return true
}

// @lc code=end
