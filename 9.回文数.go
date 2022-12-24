/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
package leetcode

import "container/list"

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	deque := list.New()

	// 先将数拆解到deque中
	for x != 0 {
		tail := x % 10
		deque.PushBack(tail)
		x = x / 10
	}

	// 两头开始取数字  分别比较
	for deque.Len() >= 2 {
		f, b := deque.Front(), deque.Back()
		fv, bv := deque.Remove(f), deque.Remove(b)
		if fv != bv {
			return false
		}
	}
	return true
}

// @lc code=end
