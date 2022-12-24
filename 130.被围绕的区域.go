/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */
package leetcode

// @lc code=start
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// if board[i][j] == 'O' {
			if i == 0 || i == m-1 {
				// 上 下
				solve_dfs(i, j, board)
			} else if j == 0 || j == n-1 {
				// 右
				solve_dfs(i, j, board)
			}
			// }
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}

}

func solve_dfs(i, j int, board [][]byte) {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
		return
	}

	if board[i][j] == 'O' {
		board[i][j] = 'A'
	} else {
		return
	}

	solve_dfs(i-1, j, board)
	solve_dfs(i+1, j, board)
	solve_dfs(i, j-1, board)
	solve_dfs(i, j+1, board)
}

// @lc code=end
