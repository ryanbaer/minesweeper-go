package lib

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

const Safe = MSValue(-2)
const Mine = MSValue(-1)
const Empty = MSValue(0)

func (m MSValue) IsMine() bool {
	return m == Mine
}

func (m MSValue) Display() rune {
	switch m {
	case Mine:
		return '*'
	case Empty:
		return ' '
	case Safe:
		return '✓'
	default:
		return rune(fmt.Sprintf("%d", m)[0])
	}
}

var (
	ErrInvalidDimensions = errors.New("Grid must be at least 2x2")
	ErrTooManyMines      = errors.New("Cannot have more mines than grid dimension")
)

type Minesweeper struct {
	Width     int
	Height    int
	Mines     int
	Remaining int
	// Grid        [][]MSValue
	Field       [][]*Location
	ActiveMines []*Location
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
		Width:       width,
		Height:      height,
		Mines:       mines,
		Remaining:   mines,
		ActiveMines: make([]*Location, 0),
	}

	m.Field = make([][]*Location, height)
	for i := 0; i < height; i++ {
		m.Field[i] = make([]*Location, width)
		for j := 0; j < width; j++ {
			m.Field[i][j] = &Location{
				Coordinate: &Coordinate{j, i},
			}
		}
	}

	// m.Grid = make([][]MSValue, height)
	// for i := 0; i < height; i++ {
	// 	m.Grid[i] = make([]MSValue, width)
	// }

	if err := m.Generate(); err != nil {
		return nil, err
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

	return nil
}

func (m *Minesweeper) Print() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			value := m.Field[i][j].Value()
			if value.IsMine() {
				fmt.Printf("*\t")
			} else {
				fmt.Printf("%d\t", value)
			}
		}
		fmt.Printf("\n\n\n")
	}
}

func (m *Minesweeper) Location(c *Coordinate) (*Location, error) {
	if !m.InBounds(c) {
		return nil, ErrOutOfBounds
	}

	return m.Field[c.Y][c.X], nil
}

func (m *Minesweeper) Perimeter(l *Location) []*Location {
	locations := make([]*Location, 0)

	for _, c := range l.RealPerimeter() {
		l1, err := m.Location(c)
		if err == nil {
			locations = append(locations, l1)
		}
	}

	return locations

}

func (m *Minesweeper) Visit(c *Coordinate) (MSValue, error) {
	l, err := m.Location(c)
	if err != nil {
		return 0, err
	}

	// m.detections()

	return l.Visit(), nil
}

func (m *Minesweeper) ValueAt(c *Coordinate) (MSValue, error) {
	l, err := m.Location(c)
	if err != nil {
		return 0, err
	}

	return l.Value(), nil
}

// func (m *Minesweeper) SetValue(c *Coordinate, value MSValue) error {
// 	if !m.InBounds(c) {
// 		return ErrOutOfBounds
// 	}
//
// 	m.Grid[c.Y][c.X] = value
// 	return nil
// }

func (m *Minesweeper) placeMines() error {
	// ct := m.Mines
	for i := 0; i < m.Mines; i++ {
		row := rand.Intn(m.Height)
		col := rand.Intn(m.Width)
		c := &Coordinate{X: col, Y: row}

		l, err := m.Location(c)
		if err != nil {
			return err
		}

		if l.Value().IsMine() {
			continue
		}
		l.setValue(Mine)
		m.ActiveMines = append(m.ActiveMines, l)
	}

	return nil
}

func (m *Minesweeper) markMines() error {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			count := MSValue(0)
			l, err := m.Location(&Coordinate{X: j, Y: i})
			if err != nil {
				return err
			}

			if !l.Value().IsMine() && l.Value() != Safe {
				for _, loc := range m.Perimeter(l) {
					if loc.Value().IsMine() {
						count += 1
					}
				}
				l.setValue(count)
			}
		}
	}

	return nil
}

func (m *Minesweeper) MarkMines() {
	m.markMines()
}

func (m *Minesweeper) Detections() {
	m.detections()
}

// detections returns
func (m *Minesweeper) detections() {
	// TODO: Remove safe mines
	// TODO: Figure out why marking safe takes an additional click after
	for _, mine := range m.ActiveMines {
		safe := true
		for _, l := range m.Perimeter(mine) {
			if !l.Visited() {
				safe = false
				break
			}
		}
		if safe {
			mine.setValue(Safe)
			mine.Visit()
		}
	}
}

func (m *Minesweeper) Size() (int, int) {
	return m.Width, m.Height
}

// Bounds is used to determine if a given coordinate is in bounds of the grid
func (m *Minesweeper) InBounds(c *Coordinate) bool {
	return !(c.X < 0 || c.Y < 0 || c.X > m.Width-1 || c.Y > m.Height-1)
}
