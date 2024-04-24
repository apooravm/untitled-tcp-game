package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

var GameBG Background

type Background struct {
	Win_Size int
	Screen   *tcell.Screen
	Style tcell.Style
}

func InitScreen() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("error creating screen.", err.Error())
		return 	
	}

	if err := screen.Init(); err != nil {
		fmt.Println("error initialising screen.", err.Error())
		return
	}

	defer screen.Fini()

	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))

	screen.Clear()

	width, height := screen.Size()
	fmt.Println(width, height)

	// BACK_STYLED := tcell.StyleDefault.
	// 	Foreground(tcell.ColorBlack).
	// 	Background(tcell.ColorWhite)
	//
	// GREY_STYLED := tcell.StyleDefault.
	// 	Foreground(tcell.ColorLightGray).
	// 	Background(tcell.ColorWhite)
	//
	// RED_STYLED := tcell.StyleDefault.
	// 	Foreground(tcell.ColorBlack).
	// 	Background(tcell.ColorRed)
	//
	// GREEN_STYLED := tcell.StyleDefault.
	// 	Foreground(tcell.ColorBlack).
	// 	Background(tcell.ColorGreen)
	//
	// addStrHoriz(0, height-1, "quit: esc", BACK_STYLED, screen)

	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			LogData(string(event.Rune())+"\n")
			if event.Key() == tcell.KeyEscape || string(event.Rune()) == "q" {
				return
			} else if event.Key() == tcell.KeyLeft {
				continue

			} else if event.Key() == tcell.KeyRight {
				continue

			} else if event.Key() == tcell.KeyUp {
				continue

			} else if event.Key() == tcell.KeyDown {
				continue

			} 		
		}
		// t1.showTime()
		screen.Show()
	}
}

func main() {
	InitScreen()
}

func LogData(data string) {
	_, err := os.ReadFile("./game.log")
	if err != nil {
		return
	}

	file, err := os.OpenFile("./game.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	defer file.Close()

	if _, err := file.Write([]byte(data)); err != nil {
		return
	}
}
