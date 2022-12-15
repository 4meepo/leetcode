/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */
package leetcode

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var high, low, highIndex int

	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i >= len(nums1) || j >= len(nums2) {
			if i >= len(nums1) {
				low = high
				high = nums2[j]
				j++
			} else if j >= len(nums2) {
				low = high
				high = nums1[i]
				i++
			}
		} else {
			if nums1[i] < nums2[j] {
				low = high
				high = nums1[i]
				i++
			} else {
				low = high
				high = nums2[j]
				j++
			}
		}

		highIndex++

		if 2*highIndex > len(nums1)+len(nums2) {
			break
		}

	}

	totalLen := len(nums1) + len(nums2)

	if totalLen%2 == 0 {
		return float64(high+low) / 2
	}
	return float64(high)

}

// @lc code=end
