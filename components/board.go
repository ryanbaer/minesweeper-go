package components

import (
	"log"

	tl "github.com/JoelOtter/termloop"
	"github.com/ryanbaer/minesweeper/lib"
)

type Board struct {
	x      int
	y      int
	width  int
	height int
	Cells  [][]*BoardCell
	ms     *lib.Minesweeper
}

func NewBoard(width, height, mines int) (*Board, error) {
	ms, err := lib.NewMinesweeper(width, height, mines)
	if err != nil {
		return nil, err
	}

	w1, h1 := ms.Size()
	cells := make([][]*BoardCell, h1)

	for i := 0; i < h1; i++ {
		row := make([]*BoardCell, w1)
		for j := 0; j < w1; j++ {
			row[j] = NewBoardCell(0, 0, '░', cellFg, cellBg)
			// level.AddEntity(row[j])
		}
		cells[i] = row
	}

	board := &Board{
		Cells:  cells,
		ms:     ms,
		width:  width,
		height: height,
	}
	// level.AddEntity(board)

	return board, nil
}

// func (b *Board) Explode() {
// 	w, h := b.ms.Size()
// 	for i := 0; i < h; i++ {
// 		for j := 0; j < w; j++ {
//
// 		}
// 	}
// }

func (b *Board) Click(tx, ty int) bool {
	var exploded bool
	if tx >= b.x && tx < b.x+b.width && ty >= b.y && ty < b.y+b.height {
		x, y := tx-b.x, ty-b.y

		val, err := b.ms.ValueAt(&lib.Coordinate{x, y})
		if err != nil {
			log.Print("Error retrieving value", err)
			return exploded
		}

		if val.IsMine() {
			exploded = true
			b.Cells[y][x].SetColor(tl.ColorRed)
			// point := lib.Coordinate{X: x, Y: y}
			// coords := point.RealPerimeter()
			// for _, c := range coords {
			// 	if b.ms.InBounds(c) {
			// 		b.Cells[c.Y][c.X].SetColor(tl.ColorRed)
			// 	}
			// }
		}

		detections := b.ms.Detections()
		for _, d := range detections {
			b.Cells[d.Y][d.X].SetContent('✓')
		}

		b.Cells[y][x].SetContent(val.Display())
	}

	return exploded
}

func (b *Board) Draw(s *tl.Screen) {
	w, h := s.Size()
	w1, h1 := b.ms.Size()
	b.x = (w/2 - w1/2)
	b.y = (h/2 - h1/2)

	for i := 0; i < h1; i++ {
		for j := 0; j < w1; j++ {
			x := b.x + j
			y := b.y + i
			cell := b.Cells[i][j]
			cell.SetPosition(x, y)
			cell.Draw(s)
		}
	}

}
