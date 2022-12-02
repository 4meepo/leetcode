/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	return middleIterator(root)
}

// 迭代实现
func middleIterator(root *TreeNode) []int {
	panic("")
}

// 递归实现
func middleRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var rtn []int
	middleTravesal(root, &rtn)

	return rtn
}

func middleTravesal(node *TreeNode, ints *[]int) {
	if node.Left != nil {
		middleTravesal(node.Left, ints)
	}
	*ints = append(*ints, node.Val)
	if node.Right != nil {
		middleTravesal(node.Right, ints)
	}
}

// @lc code=end
