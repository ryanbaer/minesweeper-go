package levels

import (
	"context"

	tl "github.com/JoelOtter/termloop"
)

const (
	titleFg = tl.ColorBlack
	titleBg = tl.ColorWhite
)

var titleContent = []string{
	"Minesweeper",
	"Press [Enter] to start",
}

func TitleLevel(ctx context.Context) MinesweeperLevel {
	level := &titleLevel{
		Content:  make([]*tl.Text, 0),
		ResultCh: make(chan MinesweeperLevel),
		ctx:      ctx,
	}
	for _, str := range titleContent {
		level.Content = append(level.Content, tl.NewText(0, 0, str, titleFg, titleBg))
	}

	level.Level = tl.NewBaseLevel(tl.Cell{
		Bg: titleBg,
		Fg: titleFg,
	})

	return level
}

type titleLevel struct {
	tl.Level
	Content  []*tl.Text
	ResultCh chan MinesweeperLevel
	ctx      context.Context
}

func (t *titleLevel) Draw(screen *tl.Screen) {
	w, h := screen.Size()
	for i, t := range t.Content {
		t.SetPosition(w/2-len(t.Text())/2, h/2+i)
		t.Draw(screen)
	}
}

func (t *titleLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey && event.Key == tl.KeyEnter {
		t.ResultCh <- MainLevel(t.Context())
	}
}

func (t *titleLevel) Result() chan MinesweeperLevel {
	return t.ResultCh
}

func (t *titleLevel) Context() context.Context {
	return t.ctx
}
