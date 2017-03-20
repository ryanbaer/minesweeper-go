package utils

import tl "github.com/JoelOtter/termloop"

// MouseUp is a helper method to return whether or not the event is a mouse up event
func MouseUp(e tl.Event) bool {
	return e.Type == tl.EventMouse && e.Key == tl.MouseRelease
}

// EnterPress is a helper method to return whether or not the event was a press of the enter key
func EnterPress(e tl.Event) bool {
	return e.Type == tl.EventKey && e.Key == tl.KeyEnter
}

// CenterText accepts a string, width, height, and row (y) and returns
// the x and y (respectively) of where the text should be positioned to
// sit in the center of the screen. It's useful for centering several rows
// of text relative to one another in the middle of the screen. `row` assumes
// a text height of 1 pixel (hence the `+ row`) at the end to detemrine the
// offset from the center
func CenterText(str string, w, h, row int) (int, int) {
	return w/2 - len(str)/2, h/2 + row
}
