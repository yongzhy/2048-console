package main

import (
	"testing"
)

type BoardTest struct {
	before [4][4]uint
	move   int
	after  [4][4]uint
}

var TestTable = []BoardTest{
    // Move LEFT
	{
		[4][4]uint{
			{0, 0, 2, 4},
			{2, 2, 4, 4},
			{2, 4, 4, 8},
			{0, 4, 4, 4}},
		MOVE_LEFT,
		[4][4]uint{
			{2, 4, 0, 0},
			{4, 8, 0, 0},
			{2, 8, 8, 0},
			{8, 4, 0, 0}},
	},

	{
		[4][4]uint{
			{2, 2, 4, 8},
			{2, 2, 4, 4},
			{0, 4, 4, 8},
			{4, 4, 8, 16}},
		MOVE_LEFT,
		[4][4]uint{
			{4, 4, 8, 0},
			{4, 8, 0, 0},
			{8, 8, 0, 0},
			{8, 8, 16, 0}},
	},

    // Move Right
	{
		[4][4]uint{
			{2, 4, 0, 0},
			{2, 2, 4, 4},
			{2, 4, 4, 8},
			{0, 4, 4, 4}},
		MOVE_RIGHT,
		[4][4]uint{
			{0, 0, 2, 4},
			{0, 0, 4, 8},
			{0, 2, 8, 8},
			{0, 0, 4, 8}},
	},

	
	{
		[4][4]uint{
			{2, 2, 4, 8},
			{2, 2, 4, 4},
			{0, 4, 4, 8},
			{4, 4, 8, 16}},
		MOVE_RIGHT,
		[4][4]uint{
			{0, 4, 4, 8},
			{0, 0, 4, 8},
			{0, 0, 8, 8},
			{0, 8, 8, 16}},
	},

    // Move Up
	{
		[4][4]uint{
			{0, 0, 2, 4},
			{2, 2, 4, 4},
			{2, 4, 4, 8},
			{0, 4, 4, 4}},
		MOVE_UP,
		[4][4]uint{
			{4, 2, 2, 8},
			{0, 8, 8, 8},
			{0, 0, 4, 4},
			{0, 0, 0, 0}},
	},

	
	{
		[4][4]uint{
			{2, 2, 4, 8},
			{2, 2, 4, 4},
			{0, 4, 4, 8},
			{4, 4, 8, 16}},
		MOVE_UP,
		[4][4]uint{
			{4, 4, 8, 8},
			{4, 8, 4, 4},
			{0, 0, 8, 8},
			{0, 0, 0, 16}},
	},

    // Move Down
	{
		[4][4]uint{
			{0, 0, 2, 4},
			{2, 2, 4, 4},
			{2, 4, 4, 8},
			{0, 4, 4, 4}},
		MOVE_DOWN,
		[4][4]uint{
			{0, 0, 0, 0},
			{0, 0, 2, 8},
			{0, 2, 4, 8},
			{4, 8, 8, 4}},
	},

	
	{
		[4][4]uint{
			{2, 2, 4, 8},
			{2, 2, 4, 4},
			{0, 4, 4, 8},
			{4, 4, 8, 16}},
		MOVE_DOWN,
		[4][4]uint{
			{0, 0, 0, 8},
			{0, 0, 4, 4},
			{4, 4, 8, 8},
			{4, 8, 8, 16}},
	},

}

func InitBoard(data [4][4]uint) {
	for i := 0; i < BOARD_ROWS; i++ {
		for j := 0; j < BOARD_COLS; j++ {
			board[i][j] = data[i][j]
		}
	}
}

func BoardEqualTo(data [4][4]uint) bool {
	for i := 0; i < BOARD_ROWS; i++ {
		for j := 0; j < BOARD_COLS; j++ {
			if board[i][j] != data[i][j] {
				return false
			}
		}
	}
	return true
}

func TestMove(t *testing.T) {
	for i := 0; i < len(TestTable); i++ {
		InitBoard(TestTable[i].before)
        switch(TestTable[i].move) {
        case MOVE_LEFT:
            moveLeft()
        case MOVE_RIGHT:
            moveRight()
        case MOVE_UP:
            moveUp()
        case MOVE_DOWN:
            moveDown()
        }
		if BoardEqualTo(TestTable[i].after) == false {
			t.Logf("Test Failure for %d case", i+1)
            t.Logf("Before Value %#v", TestTable[i].before)
            t.Logf("After  Value %#v", board)
			t.Fail()
		}
	}
}
