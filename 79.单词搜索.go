/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package leetcode

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited := [][]bool{}
				for x := 0; x < m; x++ {
					visited = append(visited, make([]bool, n))
				}
				if exist_dfs(board, visited, i, j, word, 0) {
					return true
				}
			}
		}
	}

	return false
}
func exist_dfs(board [][]byte, visited [][]bool, i, j int, word string, index int) bool {
	m, n := len(board), len(board[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return false
	}
	target := word[index]
	if board[i][j] != target {
		return false
	}
	if visited[i][j] == true {
		return false
	}
	visited[i][j] = true

	if index == len(word)-1 {
		return true
	}

	if exist_dfs(board, visited, i-1, j, word, index+1) {
		return true
	}

	if exist_dfs(board, visited, i+1, j, word, index+1) {
		return true
	}
	if exist_dfs(board, visited, i, j+1, word, index+1) {
		return true
	}

	if exist_dfs(board, visited, i, j-1, word, index+1) {
		return true
	}

	// 清理足迹
	visited[i][j] = false
	return false
}

// ["A","B","C","E"],
// ["S","F","E","S"],
// ["A","D","E","E"]
// ABCE SEEEFS
// @lc code=end
