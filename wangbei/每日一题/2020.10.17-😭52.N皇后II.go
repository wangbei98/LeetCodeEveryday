
// 1. 回溯算法

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	res := 0
	// 新建棋盘 并初始化为 全 .
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	// 判断当前位置能不能放皇后
	var isValid func(row, col int) bool
	isValid = func(row, col int) bool {
		for i := 0; i < n; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		for j := 0; j < n; j++ {
			if board[row][j] == 'Q' {
				return false
			}
		}
		i, j := row-1, col-1
		for i >= 0 && j >= 0 {
			if board[i][j] == 'Q' {
				return false
			}
			i--
			j--
		}
		i, j = row-1, col+1
		for i >= 0 && j < n {
			if board[i][j] == 'Q' {
				return false
			}
			i--
			j++
		}
		return true
	}
	// 回溯算法
	var backtrack func(row int)
	backtrack = func(row int) {
		// 如果当前行越界，则说明可以退出了
		if row == n {
			res++
			return
		}
		for col := 0; col < n; col++ {
			if !isValid(row, col) {
				continue
			}
			board[row][col] = 'Q'
			backtrack(row + 1)
			board[row][col] = '.'
		}
	}
	backtrack(0)
	return res
}