package lib

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
