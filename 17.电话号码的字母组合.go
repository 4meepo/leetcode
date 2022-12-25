/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package leetcode

import "strconv"

// @lc code=start
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	m := make(map[string]struct{})

	numberMap := [][]string{
		{},
		{},
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}

	for _, r := range digits {
		d, _ := strconv.Atoi(string(r))

		if len(m) == 0 {
			// first time
			arr := numberMap[d]
			for _, s := range arr {
				m[s] = struct{}{}
			}
			continue
		}

		_temp := []string{}
		for k := range m {
			delete(m, k)
			arr := numberMap[d]
			for _, s := range arr {
				_temp = append(_temp, k+s)
			}
		}
		for _, v := range _temp {
			m[v] = struct{}{}
		}
	}

	rtn := make([]string, 0, len(m))
	for k := range m {
		rtn = append(rtn, k)
	}
	return rtn
}

// @lc code=end
