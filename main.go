package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		width, height, mines int64
		err                  error
	)

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <width> <height> <# of mines>\n", os.Args[0])
		os.Exit(1)
	}

	width, err = strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	height, err = strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	mines, err = strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	m, err := NewMinesweeper(int(width), int(height), int(mines))
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Generate(); err != nil {
		log.Fatal(err)
	}
}
