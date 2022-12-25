/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package leetcode

import (
	"fmt"
	"testing"
)

func Test_exist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCESEEEFS"
	b := exist(board, word)
	fmt.Println(b)
}
