package components

//
// import (
// 	"fmt"
//
// 	tl "github.com/JoelOtter/termloop"
// )
//
// type MinesweeperCell struct {
// 	*Rectangle
// 	clicked bool
// }
//
// func NewMinesweeperCell(x, y, w, h int, fg tl.Attr, bg tl.Attr, char rune) *MinesweeperCell {
// 	return &MinesweeperCell{Rectangle: NewRectangle(x, y, w, h, fg, bg, char)}
// }
//
// func (c *MinesweeperCell) Tick(ev tl.Event) {
// 	x, y := c.Position()
//
// 	if ev.Type == tl.EventMouse && ev.Key == tl.MouseRelease && ev.MouseX == x && ev.MouseY == y {
// 		fmt.Printf("x: %d, y: %d, MouseX: %d, MouseY: %d\n", x, y, ev.MouseX, ev.MouseY)
// 		c.SetChar('O')
// 		// if !c.clicked {
// 		// 	c.clicked = true
// 		// 	coord := &Coordinate{
// 		// 		X: x - 1,
// 		// 		Y: y - 1,
// 		// 	}
// 		// 	val, err := ms.ValueAt(coord)
// 		// 	if err == nil {
// 		// 		if val.IsMine() {
// 		//
// 		// 		}
// 		// 		ct := fmt.Sprintf("%d", val)
// 		// 		c.SetChar(rune(ct[0]))
// 		// 	}
// 		// }
// 	}
// }
//
// func Died() {}
//
// // A type representing a 2D rectangle, with position, size and color.
// type Rectangle struct {
// 	x      int
// 	y      int
// 	width  int
// 	height int
// 	color  tl.Attr
// 	*tl.Cell
// 	clicked bool
// }
//
// // NewRectangle creates a new Rectangle at position (x, y), with size
// // (width, height) and color color.
// // Returns a pointer to the new Rectangle.
// func NewRectangle(x, y, w, h int, fg, bg tl.Attr, char rune) *Rectangle {
// 	r := Rectangle{
// 		x: x, y: y, width: w, height: h,
// 		color: bg, Cell: &tl.Cell{
// 			Bg: bg,
// 			Fg: fg,
// 			Ch: char,
// 		},
// 	}
//
// 	return &r
// }
//
// // Draws the Rectangle r onto Screen s.
// func (r *Rectangle) Draw(s *tl.Screen) {
// 	w, h := s.Size()
// 	// for i, t := range t.Content {
// 	// 	t.SetPosition(w/2-len(t.Text())/2, h/2+i)
// 	// 	t.Draw(screen)
// 	// }
//
// 	r.SetPosition(w/2-r.width/2, h/2-r.height/2)
// 	for i := 0; i <= r.height; i++ {
// 		for j := 0; j < r.width; j++ {
// 			s.RenderCell(r.x+j, r.y+i, r.Cell)
// 		}
// 	}
// }
//
// func (r *Rectangle) Tick(ev tl.Event) {}
//
// // Size returns the width and height in characters of the Rectangle.
// func (r *Rectangle) Size() (int, int) {
// 	return r.width, r.height
// }
//
// // Position returns the x and y coordinates of the Rectangle.
// func (r *Rectangle) Position() (int, int) {
// 	return r.x, r.y
// }
//
// // SetPosition sets the coordinates of the Rectangle to be x and y.
// func (r *Rectangle) SetPosition(x, y int) {
// 	r.x = x
// 	r.y = y
// }
//
// // SetSize sets the width and height of the Rectangle to be w and h.
// func (r *Rectangle) SetSize(w, h int) {
// 	r.width = w
// 	r.height = h
// }
//
// // // Color returns the color of the Rectangle.
// // func (r *Rectangle) Color() tl.Attr {
// // 	return r.color
// // }
// //
//
// func (r *Rectangle) SetChar(c rune) {
// 	r.Ch = c
// }
//
// func (r *Rectangle) SetFg(color tl.Attr) {
// 	r.Fg = color
// }
//
// func (r *Rectangle) SetBg(color tl.Attr) {
// 	r.Bg = color
// }
