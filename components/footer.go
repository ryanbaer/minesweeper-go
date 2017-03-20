package components

import tl "github.com/JoelOtter/termloop"

type Footer struct {
	fg      tl.Attr
	bg      tl.Attr
	Content []*tl.Text
	offset  int
}

func (f *Footer) Draw(screen *tl.Screen) {
	w, h := screen.Size()
	for i, t := range f.Content {
		x := w/2 - len(t.Text())/2
		y := h - (len(f.Content) - i) - f.offset

		t.SetPosition(x, y)
		t.Draw(screen)
	}
}

func (f *Footer) Append(s string) int {
	f.Content = append(f.Content, tl.NewText(0, 0, s, f.fg, f.bg))
	return len(f.Content) - 1
}

func (f *Footer) Prepend(s string) {
	f.Content = append([]*tl.Text{tl.NewText(0, 0, s, f.fg, f.bg)}, f.Content...)
}

func (f *Footer) SetContent(strs []string) {
	items := make([]*tl.Text, len(strs))
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		items[i] = tl.NewText(0, 0, str, f.fg, f.bg)
	}

	f.Content = items
	// for _, str := range c {
	// 	f.Content = append(f.Content, )
	// }
}

func (f *Footer) UpdateItem(idx int, s string) bool {
	if idx < 0 || idx > len(f.Content) {
		return false
	}

	f.Content[idx] = tl.NewText(0, 0, s, f.fg, f.bg)

	return true
}

// SetOffset adjusts how many pixels above the footer the lowest line should be
func (f *Footer) SetOffset(o int) {
	f.offset = o
}

func NewFooter(content []string, fg, bg tl.Attr, offset int) *Footer {
	f := &Footer{
		offset: offset,
		fg:     fg,
		bg:     bg,
	}

	f.SetContent(content)

	return f
}
