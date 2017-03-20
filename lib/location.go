package lib

import "fmt"

type Location struct {
	*Coordinate
	value   MSValue
	visited bool
}

func (l *Location) Visited() bool {
	return l.visited
}

func (l *Location) Visit() MSValue {
	l.visited = true
	return l.value
}

func (l *Location) Value() MSValue {
	return l.value
}

func (l *Location) setValue(value MSValue) {
	l.value = value
}

type MSValue int

const Exploded = MSValue(-4)
const Hidden = MSValue(-3)
const Safe = MSValue(-2)
const Mine = MSValue(-1)
const Empty = MSValue(0)

func (m MSValue) IsMine() bool {
	return m == Mine
}

func (m MSValue) Display() rune {
	switch m {
	case Hidden:
		return '▓'
	case Mine, Exploded:
		return '*'
	case Empty:
		// return '░'
		return ' '
	case Safe:
		return '✓'
	default:
		return rune(fmt.Sprintf("%d", m)[0])
	}
}
