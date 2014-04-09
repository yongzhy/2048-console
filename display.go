package main

import (
    "github.com/nsf/termbox-go"
    "fmt"
)

const (
    COL_OFFSET = 7
    ROW_OFFSET = 4
    ROW_COUNT = 4
    COL_COUNT = 4
    BOARD_HEIGHT = ROW_OFFSET*ROW_COUNT
    BOARD_WIDTH = COL_OFFSET*COL_COUNT

    coldef = termbox.ColorDefault
)

const GameOverMsg = "Game Over! Continue? Press [Enter]"

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
    for _, c := range msg {
        termbox.SetCell(x, y, c, fg, bg)
        x++
    }
}

func fill(x, y, w, h int, cell termbox.Cell) {
    for ly := 0; ly < h; ly++ {
        for lx := 0; lx < w; lx++ {
            termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
        }
    }
}

func drawNumber(data [4][4]uint) {
    w, h := termbox.Size()

    startx := (w - BOARD_WIDTH) / 2 
    starty := (h - BOARD_HEIGHT) / 2 

    for row:=0; row < 4; row++ {
        for col:=0; col < 4; col++ {
            x := startx + col * COL_OFFSET + 1
            y := starty + row * ROW_OFFSET + 1
            if data[row][col] > 0 {
        		tbprint(x, y, coldef, coldef, fmt.Sprintf("%4d", data[row][col]))
        	} else {
        		tbprint(x, y, coldef, coldef, "    ")
        	}
        }
    }
    termbox.Flush()
}

func drawScore(score uint) {
    w, h := termbox.Size()

    startx := (w - BOARD_WIDTH) / 2 
    starty := (h - BOARD_HEIGHT) / 2 

    tbprint(startx+1, starty-3, coldef, coldef, fmt.Sprintf("Score: %-10d", score))
    termbox.Flush()
}

func drawTitle() {
    w, h := termbox.Size()

    startx := (w - BOARD_WIDTH) / 2
    starty := (h - BOARD_HEIGHT) / 2

    tbprint(startx+1, starty-5, coldef, coldef, "========== 2048 ==========")
    termbox.Flush()    
}

func drawGameOver(clear bool) {
    w, h := termbox.Size()

    startx := (w - BOARD_WIDTH) / 2
    starty := (h - BOARD_HEIGHT) / 2	
	if clear {
		fill(startx + 1, starty+BOARD_HEIGHT+6, len(GameOverMsg), 1, termbox.Cell{Ch: ' '})   
	} else {
		tbprint(startx + 1, starty+BOARD_HEIGHT+6, coldef, coldef, GameOverMsg)  
	}
}

func drawGrid() {
    termbox.Clear(coldef, coldef)
    w, h := termbox.Size()

    startx := (w - BOARD_WIDTH) / 2
    starty := (h - BOARD_HEIGHT) / 2

    fill(startx-1, starty, 1, BOARD_HEIGHT -1, termbox.Cell{Ch: '|'})
    fill(startx-1 + COL_OFFSET, starty, 1, BOARD_HEIGHT -1, termbox.Cell{Ch: '|'})
    fill(startx-1 + COL_OFFSET*2, starty, 1, BOARD_HEIGHT -1, termbox.Cell{Ch: '|'})
    fill(startx-1 + COL_OFFSET*3, starty, 1, BOARD_HEIGHT -1, termbox.Cell{Ch: '|'})

    fill(startx, starty-1, BOARD_WIDTH-1, 1, termbox.Cell{Ch: '─'})
    fill(startx, starty-1 + ROW_OFFSET*1, BOARD_WIDTH-1, 1, termbox.Cell{Ch: '─'})
    fill(startx, starty-1 + ROW_OFFSET*2, BOARD_WIDTH-1, 1, termbox.Cell{Ch: '─'})
    fill(startx, starty-1 + ROW_OFFSET*3, BOARD_WIDTH-1, 1, termbox.Cell{Ch: '─'})
    fill(startx, starty-1 + ROW_OFFSET*4, BOARD_WIDTH-1, 1, termbox.Cell{Ch: '─'})

    fill(startx-1 + COL_OFFSET*4, starty, 1, BOARD_HEIGHT -1, termbox.Cell{Ch: '|'})

    termbox.SetCell(startx-1, starty-1, '┌', coldef, coldef)
    termbox.SetCell(startx-1, starty+ROW_OFFSET*4 -1, '└', coldef, coldef)
    termbox.SetCell(startx+COL_OFFSET*4-1, starty-1, '┐', coldef, coldef)
    termbox.SetCell(startx+COL_OFFSET*4-1, starty+ROW_OFFSET*4-1, '┘', coldef, coldef)

    tbprint(startx + 1, starty+BOARD_HEIGHT+3, coldef, coldef, "Press ESC to quit")
    termbox.Flush()
}


    
