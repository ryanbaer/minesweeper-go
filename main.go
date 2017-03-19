package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ryanbaer/minesweeper/levels"
)

//
// var (
// 	ms *Minesweeper
// )

func main() {
	var (
		width, height, mines int
		err                  error
	)

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <width> <height> <# of mines>\n", os.Args[0])
		os.Exit(1)
	}

	w, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	h, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	m, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	width = int(w)
	height = int(h)
	mines = int(m)

	levels.StartGame(&levels.Config{
		Width:  width,
		Height: height,
		Mines:  mines,
		TitleContent: []string{
			"Minesweeper",
			"Press [enter] to start",
		},
		WinContent:  []string{"Congratulations!!", "You did it!"},
		LoseContent: []string{"You exploded!"},
		QuitMessage: "Press Ctrl + C to quit",
	})
}
