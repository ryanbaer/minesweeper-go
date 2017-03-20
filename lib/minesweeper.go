package lib

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Seed random for mine placement
func init() {
	rand.Seed(time.Now().Unix())
}

// DEBUG: Write to logfile
func init() {
	// logfile := "minesweeper.log"
	// if _, err := os.Stat(logfile); !os.IsNotExist(err) {
	// 	os.Remove(logfile)
	// }
	//
	// f, err := os.Create(logfile)
	// if err == nil {
	// 	log.SetOutput(f)
	// } else {
	// 	log.Print("Unable to write to log file")
	// }
}

const MinWidth = 2
const MinHeight = 2

var (
	ErrInvalidDimensions = errors.New("Grid must be at least 2x2")
	ErrTooManyMines      = errors.New("Cannot have grid of more than half mines")
)

type Minesweeper struct {
	Width       int
	Height      int
	Mines       int
	Field       [][]*Location
	ActiveMines []*Location
}

func NewMinesweeper(width, height, mines int) (*Minesweeper, error) {
	if width < MinWidth || height < MinHeight {
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

	if err := m.Generate(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Minesweeper) Remaining() int {
	return len(m.ActiveMines)
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

func (m *Minesweeper) DisplayGrid() []string {
	rows := make([]string, 0)
	for i := 0; i < m.Height; i++ {
		var str = ""
		for j := 0; j < m.Width; j++ {
			str += fmt.Sprintf("%c", m.Field[i][j].Value().Display())
		}
		rows = append(rows, str)
	}
	return rows
}

func (m *Minesweeper) Print() {
	for _, row := range m.DisplayGrid() {
		log.Print(row)
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

	// TODO: Figure out why marking safe here in detections takes an additional click after
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

func (m *Minesweeper) placeMines() error {
	remaining := m.Mines
	for remaining > 0 {
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
		remaining--
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

			// TODO: Do better than not safe && not mine. maybe a test that combines them?
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

func (m *Minesweeper) Detections() bool {
	return m.detections()
}

// detections iterates the list of active mines and finds any that are safe,
// removing them from the list. It returns a bool, that when true means there
// are no remaining active mines and the game has been won
func (m *Minesweeper) detections() bool {
	var remove = make([]*Location, 0)
	for _, mine := range m.ActiveMines {
		safe := true
		for _, l := range m.Perimeter(mine) {
			if !l.Value().IsMine() && !l.Visited() { // TODO: Bug: if another mine is in permeter, isn't solved - could just do !mine?
				safe = false
				break
			}
		}

		if safe {
			// Mark mine safe, visited, and mark it for removal from the active mine list
			mine.setValue(Safe)
			mine.Visit()
			remove = append(remove, mine)
		}
	}

	for _, mine := range remove {
		if idx := findLocation(mine, m.ActiveMines); idx != -1 {
			m.ActiveMines = removeLocation(idx, m.ActiveMines)
		}
	}

	return len(m.ActiveMines) == 0
}

func findLocation(l *Location, slice []*Location) int {
	for i, loc := range slice {
		if l == loc {
			return i
		}
	}

	return -1
}

func removeLocation(idx int, slice []*Location) []*Location {
	copy(slice[idx:], slice[idx+1:])
	slice[len(slice)-1] = nil // or the zero value of T
	slice = slice[:len(slice)-1]

	return slice
}

func (m *Minesweeper) Size() (int, int) {
	return m.Width, m.Height
}

// Bounds is used to determine if a given coordinate is in bounds of the grid
func (m *Minesweeper) InBounds(c *Coordinate) bool {
	return !(c.X < 0 || c.Y < 0 || c.X > m.Width-1 || c.Y > m.Height-1)
}
