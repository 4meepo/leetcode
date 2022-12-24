/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package leetcode

import (
	"fmt"
	"testing"
)

func Test_numIslands(t *testing.T) {
	fmt.Println('1')
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	n := numIslands(grid)
	fmt.Println(n)

	grid = [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	n = numIslands(grid)
	fmt.Println(n)
}
