package components

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

const cellBg = tl.ColorWhite
const cellFg = tl.ColorBlack

type BoardCell struct {
	*tl.Rectangle
	Content rune
	bg      tl.Attr
	fg      tl.Attr
}

func NewBoardCell(x, y int, content rune, fg, bg tl.Attr) *BoardCell {
	rect := tl.NewRectangle(x, y, 1, 1, fg)
	return &BoardCell{
		Rectangle: rect,
		Content:   content,
	}
}

func (b *BoardCell) SetColor(color tl.Attr) {
	b.bg = color
	b.fg = color
}

func (b *BoardCell) Draw(s *tl.Screen) {
	x, y := b.Position()
	// fmt.Printf("x: %d, y: %d\n", x, y)
	s.RenderCell(x, y, &tl.Cell{Bg: b.bg, Fg: b.fg, Ch: b.Content})
}

func (b *BoardCell) SetContent(content rune) {
	b.Content = content
}

func (b *BoardCell) Tick(ev tl.Event) {
	x, y := b.Position()
	fmt.Println(x, y)
	if ev.Type == tl.EventMouse && ev.Key == tl.MouseLeft && ev.MouseX == x && ev.MouseY == y {
		// fmt.Printf("x: %d, y: %d, MouseX: %d, MouseY: %d\n", x, y, ev.MouseX, ev.MouseY)
		b.SetContent('O')
	}
}
