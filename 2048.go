package main 

import (
    "github.com/nsf/termbox-go"
)

func main() {
    
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    termbox.HideCursor()
    
    termbox.SetInputMode(termbox.InputEsc )

	initBoard()
    drawGrid()
    drawTitle()
    drawNumber(board)
    drawScore(score)

    gameOver := false
    move := MOVE_NONE
mainloop:
    for {
    	move = MOVE_NONE
        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            switch ev.Key {
            case termbox.KeyEsc:
                break mainloop
            case termbox.KeyArrowLeft:
            	move = MOVE_LEFT
            case termbox.KeyArrowRight:
            	move = MOVE_RIGHT
            case termbox.KeyArrowUp:
            	move = MOVE_UP
            case termbox.KeyArrowDown:
            	move = MOVE_DOWN
            case termbox.KeyEnter:
            	if gameOver {
            		gameOver = false
            		initBoard()
            		drawGameOver(true)
            	}
            default:
            }
        case termbox.EventError:
            panic(ev.Err)
        }

        valid := doMove(move)
        gameOver = isGameOver()
        if gameOver == true {
			drawGameOver(false)
		} else if valid == true {
			addNumber(move)
		}

    	drawNumber(board)
    	drawScore(score)
    }    
}