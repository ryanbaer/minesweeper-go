package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Seed random for mine placement
func init() {
	rand.Seed(time.Now().Unix())
}

// TODO: Integrate Coordinates into grid so you can just have:
// func (m *Minesweeper) Coordinate(x, y int) (*Coordinate, error)
// and you can simply do the bounds check in there because
// Coordinate determines whether or not the pair fall in the grid

type MSValue int

const minWidth = 2
const minHeight = 2

const Mine = MSValue(-1)

func (m MSValue) IsMine() bool {
	return m == Mine
}

var (
	ErrInvalidDimensions = errors.New("Grid must be at least 2x2")
	ErrTooManyMines      = errors.New("Cannot have more mines than grid dimension")
)

type Minesweeper struct {
	Width  int
	Height int
	Mines  int
	Grid   [][]MSValue
}

func NewMinesweeper(width, height, mines int) (*Minesweeper, error) {
	if width < minWidth || height < minHeight {
		return nil, ErrInvalidDimensions
	}

	maxMines := width*height - 1
	if mines > maxMines {
		return nil, ErrTooManyMines
	}

	m := &Minesweeper{
		Width:  width,
		Height: height,
		Mines:  mines,
	}

	m.Grid = make([][]MSValue, height)
	for i := 0; i < height; i++ {
		m.Grid[i] = make([]MSValue, width)
	}

	return m, nil
}

func (m *Minesweeper) Generate() error {
	if err := m.placeMines(); err != nil {
		return err
	}

	if err := m.markMines(); err != nil {
		return err
	}

	m.Print()

	return nil
}

func (m *Minesweeper) Print() {
	for i := 0; i < len(m.Grid); i++ {
		for j := 0; j < len(m.Grid[0]); j++ {
			val := m.Grid[i][j]
			if val == -1 {
				fmt.Printf("*\t")
			} else {
				fmt.Printf("%d\t", m.Grid[i][j])
			}
		}
		fmt.Printf("\n\n\n")
	}
}

func (m *Minesweeper) ValueAt(c *Coordinate) (MSValue, error) {
	if !m.Bounds(c) {
		return 0, ErrOutOfBounds
	}
	return m.Grid[c.Y][c.X], nil
}

func (m *Minesweeper) SetValue(c *Coordinate, value MSValue) error {
	if !m.Bounds(c) {
		return ErrOutOfBounds
	}

	m.Grid[c.Y][c.X] = value
	return nil
}

func (m *Minesweeper) placeMines() error {
	ct := m.Mines
	for ct > 0 {
		row := rand.Intn(len(m.Grid))
		col := rand.Intn(len(m.Grid[0]))
		c := &Coordinate{X: col, Y: row}

		val, err := m.ValueAt(c)
		if err != nil {
			return err
		}

		if val == Mine {
			continue
		}

		if err := m.SetValue(c, Mine); err != nil {
			return err
		}
		ct--
	}

	return nil
}

func (m *Minesweeper) markMines() error {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			count := MSValue(0)
			c := &Coordinate{X: j, Y: i}

			val, err := m.ValueAt(c)
			if err != nil {
				return err
			}

			if !val.IsMine() {
				for _, offset := range c.Perimeter() {
					check := c.Add(offset)
					val1, err := m.ValueAt(check)
					if err != nil {
						if err != ErrOutOfBounds {
							return err
						}
						continue
					}

					if val1.IsMine() {
						count += 1
					}
				}

				err := m.SetValue(c, count)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Bounds is used to determine if a given coordinate is in bounds of the grid
func (m *Minesweeper) Bounds(c *Coordinate) bool {
	return !(c.X < 0 || c.Y < 0 || c.X > m.Width-1 || c.Y > m.Height-1)
}
