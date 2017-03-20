package components

import (
	"log"

	tl "github.com/JoelOtter/termloop"
	"github.com/ryanbaer/minesweeper-go/lib"
)

const cellBg = tl.ColorWhite
const cellFg = tl.ColorBlack

// const cellBlank = '░'
// const cellHidden = '▓'
const cellHidden = '░'

type Board struct {
	x      int
	y      int
	width  int
	height int
	cells  [][]*tl.Cell
	shown  bool
	*lib.Minesweeper
}

func NewBoard(width, height, mines int) (*Board, error) {
	ms, err := lib.NewMinesweeper(width, height, mines)
	if err != nil {
		return nil, err
	}

	w1, h1 := ms.Size()
	cells := make([][]*tl.Cell, h1)

	for i := 0; i < h1; i++ {
		row := make([]*tl.Cell, w1)
		for j := 0; j < w1; j++ {
			row[j] = &tl.Cell{Fg: cellFg, Bg: cellBg}
		}

		cells[i] = row
	}

	board := &Board{
		cells:       cells,
		Minesweeper: ms,
		// ms:     ms,
		width:  width,
		height: height,
	}

	return board, nil
}

// func (b *Board) Explode() {
// 	w, h := b.Size()
// 	for i := 0; i < h; i++ {
// 		for j := 0; j < w; j++ {
//
// 		}
// 	}
// }

const (
	_ = iota
	GameStateActive
	GameStateWon
	GameStateLost
)

type GameState int

func (b *Board) ToggleSolution() bool {
	b.shown = !b.shown
	return b.shown
}

func (b *Board) Click(tx, ty int) GameState {
	var (
		exploded bool
	)
	// log.Printf("Click: %d, %d\nRect: (%d, %d) (%d, %d)", tx, ty, b.x, b.y, b.x+b.width, b.y+b.height)
	if tx >= b.x && tx < b.x+b.width && ty >= b.y && ty < b.y+b.height {
		// log.Printf("Click: %d, %d in bounds", tx, ty)
		x, y := tx-b.x, ty-b.y
		b.cells[y][x].Bg = cellBg
		value, err := b.Visit(&lib.Coordinate{x, y})
		if err != nil {
			log.Print(err)
			return -1
		}

		exploded = value.IsMine()

		if exploded {
			return GameStateLost
		}

		if won := b.Detections(); won {
			return GameStateWon
		}
		b.MarkMines()
	}

	return GameStateActive
}

// func (b *Board) Press(tx, ty int) {
//
// 	if tx >= b.x && tx < b.x+b.width && ty >= b.y && ty < b.y+b.height {
// 		x, y := tx-b.x, ty-b.y
// 		b.cells[y][x].Bg = tl.ColorBlue
// 	}
// }

func (b *Board) Draw(s *tl.Screen) {
	w, h := s.Size()
	w1, h1 := b.Size()
	b.x = (w/2 - w1/2)
	b.y = (h/2 - h1/2)

	for i := 0; i < h1; i++ {
		for j := 0; j < w1; j++ {
			var value rune
			x := b.x + j
			y := b.y + i

			loc, err := b.Location(&lib.Coordinate{j, i})
			if err != nil {
				log.Print(err)
				continue
			}

			value = loc.Value().Display()

			if !b.shown && !loc.Visited() {
				value = cellHidden
			}
			// } else {
			// 	if loc.Visited() {
			// 		value = loc.Value().Display()
			// 	} else {
			//
			// 	}
			// }

			cell := b.cells[i][j]
			cell.Ch = value
			s.RenderCell(x, y, cell)
		}
	}

}
