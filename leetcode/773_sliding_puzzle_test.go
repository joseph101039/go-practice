package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlidingPuzzle(t *testing.T) {
	assert.Equal(t, 0, slidingPuzzle([][]int{{1, 2, 3}, {4, 5, 0}}))
	assert.Equal(t, 1, slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
	assert.Equal(t, -1, slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}}))
	assert.Equal(t, 5, slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}}))
}

const x, y = 2, 3

type Node struct {
	board        [x][y]int8
	deep         int8
	zeroX, zeroY int8 // 當前 0 的位置
}

func slidingPuzzle(board [][]int) (step int) {
	boardInput := [x][y]int8{}
	for i, row := range board {
		for j, val := range row {
			boardInput[i][j] = int8(val)
		}
	}
	root := Node{[x][y]int8{{1, 2, 3}, {4, 5, 0}}, 0, 1, 2}
	boardMap := map[[x][y]int8]*Node{root.board: &root}

	var (
		queue []Node
		cur   = root
	)
	for {
		// step up
		if cur.zeroX-1 >= 0 {
			next := cur
			next.board[cur.zeroX][cur.zeroY], next.board[cur.zeroX-1][cur.zeroY] = next.board[cur.zeroX-1][cur.zeroY], next.board[cur.zeroX][cur.zeroY]
			if _, ok := boardMap[next.board]; !ok {
				next.deep++
				next.zeroX--
				queue = append(queue, next)
				boardMap[next.board] = &next
			}
		}

		// step down
		if cur.zeroX+1 < x {
			next := cur
			next.board[cur.zeroX][cur.zeroY], next.board[cur.zeroX+1][cur.zeroY] = next.board[cur.zeroX+1][cur.zeroY], next.board[cur.zeroX][cur.zeroY]
			if _, ok := boardMap[next.board]; !ok {
				next.deep++
				next.zeroX++
				queue = append(queue, next)
				boardMap[next.board] = &next
			}
		}

		// step left
		if cur.zeroY-1 >= 0 {
			next := cur
			next.board[cur.zeroX][cur.zeroY], next.board[cur.zeroX][cur.zeroY-1] = next.board[cur.zeroX][cur.zeroY-1], next.board[cur.zeroX][cur.zeroY]
			if _, ok := boardMap[next.board]; !ok {
				next.deep++
				next.zeroY--
				queue = append(queue, next)
				boardMap[next.board] = &next
			}
		}

		// step right
		if cur.zeroY+1 < y {
			next := cur
			next.board[cur.zeroX][cur.zeroY], next.board[cur.zeroX][cur.zeroY+1] = next.board[cur.zeroX][cur.zeroY+1], next.board[cur.zeroX][cur.zeroY]
			if _, ok := boardMap[next.board]; !ok {
				next.deep++
				next.zeroY++
				queue = append(queue, next)
				boardMap[next.board] = &next
			}
		}

		if boardMap[boardInput] != nil {
			return int(boardMap[boardInput].deep)
		}

		if len(queue) == 0 {
			break
		}

		cur, queue = queue[0], queue[1:]
	}
	return -1
}
