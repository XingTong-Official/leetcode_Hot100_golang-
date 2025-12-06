package main

func main() {
	exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")
}
func exist(board [][]byte, word string) bool {
	m := make([][]bool, len(board))
	for i := range m {
		m[i] = make([]bool, len(board[0]))
	}
	answer := ""
	path := []byte{}
	var dfs func(i, j int, index int)
	dfs = func(i, j, index int) {
		if index == len(word) {
			answer = string(path)
			return
		}
		if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || m[i][j] || board[i][j] != word[index] {
			return
		}
		m[i][j] = true
		path = append(path, board[i][j])
		dfs(i, j+1, index+1)
		dfs(i+1, j, index+1)
		dfs(i-1, j, index+1)
		dfs(i, j-1, index+1)
		m[i][j] = false
		path = path[:len(path)-1]
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				dfs(i, j, 0)
			}
		}
	}
	return answer != ""
}
