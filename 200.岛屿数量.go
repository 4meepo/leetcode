/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package leetcode

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				numIslands_dfs(grid, i, j)
			}
		}
	}

	return count
}

func numIslands_dfs(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if grid[i][j] == '1' {
		// mark visited
		grid[i][j] = '2'
		// down
		numIslands_dfs(grid, i+1, j)
		// up
		numIslands_dfs(grid, i-1, j)
		// right
		numIslands_dfs(grid, i, j+1)
		// left
		numIslands_dfs(grid, i, j-1)
	}

}

// @lc code=end
