/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
package leetcode

import (
	"fmt"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	_1 := &TreeNode{Val: 1}
	_2 := &TreeNode{Val: 2}
	_2_ := &TreeNode{Val: 2}
	_3 := &TreeNode{Val: 3}
	_4 := &TreeNode{Val: 4}
	_3_ := &TreeNode{Val: 3}
	_4_ := &TreeNode{Val: 4}

	_1.Left = _2
	_1.Right = _2_
	_2.Left = _3
	_2.Right = _4
	_2_.Left = _4_
	_2_.Right = _3_

	fmt.Println(isSymmetric(_1))

}
