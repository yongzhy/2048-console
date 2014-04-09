package main

import (
	"math/rand"
)

const (
	CLEAR      = 0
	OK = 1
	BOARD_ROWS = 4
	BOARD_COLS = 4
)

const (
	MOVE_NONE = 0
	MOVE_LEFT  = 1
	MOVE_RIGHT = 2
	MOVE_UP    = 3
	MOVE_DOWN  = 4
)

const (
	SCORE_ADD_NUMBER = 4
)

var board [4][4]uint
var score uint

func initBoard() {
	for i := 0; i < BOARD_ROWS; i++ {
		for j := 0; j < BOARD_COLS; j++ {
			board[i][j] = CLEAR
		}
	}

	row := rand.Intn(BOARD_ROWS)
	col := rand.Intn(BOARD_COLS)
	board[row][col] = 2

	score = 0
}

type GameLine struct {
	data  []uint
	count int
}

func newGameLine() *GameLine {
    return &GameLine{make([]uint, 4), 0}
}

func (b *GameLine) push(a uint) {
	if a != CLEAR {
		if b.count < len(b.data) {
			b.data[b.count] = a
			b.count++
		}
	}
}

func (b *GameLine) pop() (a uint) {
	a = 0
    offset := 1
    if b.count > 0 {
        a = b.data[0]
        if b.count >= 2 {
            if b.data[0] == b.data[1] {
                a = b.data[0] << 1
                offset = 2
                score += a
            }
        }
		for i := 0; i < b.count-offset; i++ {
			b.data[i] = b.data[i+offset]
		}
		b.count = b.count - offset
	}
	return
}

func (b *GameLine) size() int {
    return b.count
}

func (b *GameLine) clear() {
	b.count = 0
}

func addNumber(move int) {
	var number uint = 2
	if rand.Intn(1000) > 700 {
		number = 4
	}

	moveon := false
	switch move {
	case MOVE_LEFT:
		for row:=0; row<BOARD_ROWS;row++ {
			if board[row][3] & 0xFFFF == 0 {
				moveon = true
				break
			}
		}
	case MOVE_RIGHT:
		for row:=0; row<BOARD_ROWS;row++ {
			if board[row][0] & 0xFFFF == 0 {
				moveon = true
				break
			}
		}		
	case MOVE_UP:
		for col:=0; col<BOARD_COLS;col++ {
			if board[3][col] & 0xFFFF == 0 {
				moveon = true
				break
			}
		}		
	case MOVE_DOWN:
		for col:=0; col<BOARD_COLS;col++ {
			if board[0][col] & 0xFFFF == 0 {
				moveon = true
				break
			}
		}		
	}

	for moveon==true {
		a := rand.Intn(4)

		switch move {
		case MOVE_LEFT:
			if board[a][3] == 0 {
				board[a][3] = number
				score += SCORE_ADD_NUMBER
				moveon = false
			}
		case MOVE_RIGHT:
			if board[a][0] == 0 {
				board[a][0] = number
				score += SCORE_ADD_NUMBER
				moveon = false
			}			
		case MOVE_UP:
			if board[3][a] == 0 {
				board[3][a] = number
				score += SCORE_ADD_NUMBER
				moveon = false
			}			
		case MOVE_DOWN:
			if board[0][a] == 0 {
				board[0][a] = number
				score += SCORE_ADD_NUMBER
				moveon = false
			}	
		}
	}
}

func isGameOver() bool {
	// Any cell with 0 value indicate game can continue
	for row:=0; row<BOARD_ROWS; row++ {
		for col:=0; col<BOARD_COLS; col++ {
			if board[row][col] == CLEAR {
				return false
			}
		}
	}

	// any two horizental adjacent cell with same value indicate game can continue
	for row:=0; row<BOARD_ROWS; row++ {
		for col:=1; col<BOARD_COLS; col++ {
			if board[row][col] == board[row][col-1] {
				return false
			}
		}
	}

	// any two vertical adjacent cell with same value indicate game can continue
	for col:=0; col<BOARD_COLS; col++ {
		for row:=1; row<BOARD_ROWS; row++ {
			if board[row][col] == board[row-1][col] {
				return false
			}
		}
	}
	return true
}

func moveLeft() (valid bool) {
	valid = false
    queue := newGameLine()

	for row := 0; row < BOARD_ROWS; row++ {
		queue.clear()
		for col := 0; col < BOARD_COLS; col++ {
			queue.push(board[row][col])
		}
		for col := 0; col < BOARD_COLS; col++ {
			val := queue.pop()
			if board[row][col] != val {
				board[row][col] = val
				valid = true
			}
		}
	}
	return
}

func moveRight() (valid bool) {
	valid = false
    queue := newGameLine()

	for row := 0; row < BOARD_ROWS; row++ {
		queue.clear()
		for col := BOARD_COLS-1; col >= 0; col-- {
			queue.push(board[row][col])
		}
		for col := BOARD_COLS-1; col >= 0; col-- {
			val := queue.pop()
			if board[row][col] != val {
				board[row][col] = val
				valid = true
			}
		}
	}
	return
}

func moveUp() (valid bool) {
	valid = false
    queue := newGameLine()

	for col := 0; col < BOARD_COLS; col++ {
		queue.clear()
		for row := 0; row < BOARD_ROWS; row++ {
			queue.push(board[row][col])
		}
		for row := 0; row < BOARD_ROWS; row++ {
			val := queue.pop()
			if board[row][col] != val {
				board[row][col] = val
				valid = true
			}
		}
	}
	return
}

func moveDown() (valid bool) {
	valid = false
    queue := newGameLine()

	for col := 0; col < BOARD_COLS; col++ {
		queue.clear()
		for row := BOARD_ROWS-1; row >= 0; row-- {
			queue.push(board[row][col])
		}
		for row := BOARD_ROWS-1; row >= 0; row-- {
			val := queue.pop()
			if board[row][col] != val {
				board[row][col] = val
				valid = true
			}
		}
	}
	return
}

func doMove(move int) (valid bool) {
	switch(move) {
	case MOVE_LEFT:
		valid = moveLeft()
	case MOVE_RIGHT:
		valid = moveRight()
	case MOVE_UP:
		valid = moveUp()
	case MOVE_DOWN:
		valid = moveDown()
	}
	return
}