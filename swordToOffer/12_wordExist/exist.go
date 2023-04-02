package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	bs := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == bs[0] {
				if correct(board, i, j, bs, 0) {
					return true
				}
			}
		}
	}
	return false
}

func correct(board [][]byte, i int, j int, word []byte, now int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || word[now] != board[i][j] {
		return false
	}
	if now == len(word)-1 {
		return true
	}
	board[i][j] = ' '
	res := correct(board, i-1, j, word, now+1) || correct(board, i+1, j, word, now+1) ||
		correct(board, i, j-1, word, now+1) || correct(board, i, j+1, word, now+1)
	board[i][j] = word[now]
	return res
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCCED"))
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCB"))
}
