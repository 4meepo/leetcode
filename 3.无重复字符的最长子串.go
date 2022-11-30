/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

import "math"

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	longest := 0
	window := make([]rune, 0)

	for _, c := range s {
		for j := 0; j < len(window); j++ {
			if window[j] == c {
				if j+1 <= len(window) {
					// 非连续重复
					window = window[j+1:]
				} else {
					// 连续两个字符重复
					window = []rune{}
				}
				break
			}
		}
		window = append(window, c)
		longest = int(math.Max(float64(len(window)), float64(longest)))
	}

	return longest
}

// @lc code=end
