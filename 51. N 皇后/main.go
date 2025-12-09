package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}
func solveNQueens(n int) [][]string {
	answer := [][]string{}
	chessboard := make([][]byte, n)

	//init ChessBoard
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '_'
		}
		chessboard[i] = append([]byte{}, row...)
	}

	var NQueens func(index int)
	NQueens = func(index int) {
		if index == n {
			//处理byte类型的棋盘,修改所有的内容数字为'.',追加到answer里面
			temp := []string{}
			for _, row := range chessboard {
				temp2 := append([]byte{}, row...)
				for i, _ := range temp2 {
					if temp2[i] <= '9' && temp2[i] >= '0' {
						temp2[i] = '.'
					}
				}
				temp = append(temp, string(temp2))
			}
			answer = append(answer, temp)
			return
		}
		flag := 0
		for i := 0; i < n; i++ {
			if chessboard[index][i] == '_' {
				flag = 1
				chessboard[index][i] = 'Q'
				//行禁区设置
				for iRow, _ := range chessboard[index] {
					if iRow == i {
						continue
					}
					if chessboard[index][iRow] == '_' {
						chessboard[index][iRow] = byte('0' + index)
					}

				}
				//列禁区设置
				for j := 0; j < n; j++ {
					if j == index {
						continue
					}
					if chessboard[j][i] == '_' {
						chessboard[j][i] = byte('0' + index)
					}
				}
				//45°方向禁区设置
				x, y := i, index
				for {
					x++
					y--
					if y < 0 || x >= n {
						break
					}
					if chessboard[y][x] == '_' {
						chessboard[y][x] = byte('0' + index)
					}

				}
				//135°方向禁区设置
				x, y = i, index
				for {
					x++
					y++
					if y >= n || x >= n {
						break
					}
					if chessboard[y][x] == '_' {
						chessboard[y][x] = byte('0' + index)
					}
				}
				//-135°方向禁区设置
				x, y = i, index
				for {
					x--
					y++
					if y >= n || x < 0 {
						break
					}
					if chessboard[y][x] == '_' {
						chessboard[y][x] = byte('0' + index)
					}
				}
				//-45°方向禁区设置
				x, y = i, index
				for {
					x--
					y--
					if y < 0 || x < n {
						break
					}
					if chessboard[y][x] == '_' {
						chessboard[y][x] = byte('0' + index)
					}
				}
				NQueens(index + 1)
				//回溯
				chessboard[index][i] = '_'
				for iRow, _ := range chessboard[index] {
					if chessboard[index][iRow] == byte('0'+index) {
						chessboard[index][iRow] = '_'
					}

				}
				for j := 0; j < n; j++ {
					if chessboard[j][i] == byte('0'+index) {
						chessboard[j][i] = '_'
					}
				}
				//45°方向禁区复位
				x, y = i, index
				for {
					x++
					y--
					if y < 0 || x >= n {
						break
					}
					if chessboard[y][x] == byte('0'+index) {
						chessboard[y][x] = '_'
					}
				}
				//135°方向禁区复位
				x, y = i, index
				for {
					x++
					y++
					if y >= n || x >= n {
						break
					}
					if chessboard[y][x] == byte('0'+index) {
						chessboard[y][x] = '_'
					}
				}
				//-135°方向禁区复位
				x, y = i, index
				for {
					x--
					y++
					if y >= n || x < 0 {
						break
					}
					if chessboard[y][x] == byte('0'+index) {
						chessboard[y][x] = '_'
					}
				}
				//-45°方向禁区复位
				x, y = i, index
				for {
					x--
					y--
					if y < 0 || x < n {
						break
					}
					if chessboard[y][x] == byte('0'+index) {
						chessboard[y][x] = '_'
					}
				}
			}

		}
		if flag == 0 {
			return
		}
	}
	NQueens(0)

	return answer
}
