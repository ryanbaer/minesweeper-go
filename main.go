package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	tl "github.com/JoelOtter/termloop"
	"github.com/ryanbaer/minesweeper-go/levels"
)

const (
	defaultWidth  = 20
	defaultHeight = 10
	defaultMines  = 10
)

func usage(prog string) {
	fmt.Printf("\tUsage: %s <width> <height> <# of mines>\n", prog)
	fmt.Printf("\tDefault: %s %d %d %d\n", prog, defaultWidth, defaultHeight, defaultMines)
	os.Exit(1)
}

func main() {
	var (
		width, height, mines int
		err                  error
	)

	helpPtr := flag.Bool("help", false, "Print usage")
	flag.Parse()
	if *helpPtr {
		usage(os.Args[0])
	}

	if len(os.Args) == 4 {
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
	} else {
		width = defaultWidth
		height = defaultHeight
		mines = defaultMines
	}

	err = levels.StartGame(&levels.Config{
		Width:  width,
		Height: height,
		Mines:  mines,
		TitleContent: []string{
			"Minesweeper",
			"Press [enter] to start",
		},
		WinContent: []string{
			"Congratulations!!",
			"A winner is you!",
			"",
			"Press [enter] to play again",
			levels.QuitMessage,
		},
		LoseContent: []string{
			"Better luck next time!",
			"",
			"Press [enter] to try again",
			levels.QuitMessage,
		},
		FgColor: tl.ColorBlack,
		BgColor: tl.ColorWhite,
	})
	if err != nil {
		log.Fatal(err)
	}

}
