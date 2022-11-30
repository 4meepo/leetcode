/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int)

	for i, v := range nums {
		if index, ok := nmap[target-v]; ok {
			return []int{index, i}
		} else {
			nmap[v] = i
		}
	}

	return []int{}
}

// @lc code=end
