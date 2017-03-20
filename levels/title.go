package levels

import (
	"context"
	"log"

	tl "github.com/JoelOtter/termloop"
	"github.com/ryanbaer/minesweeper-go/utils"
)

type titleLevel struct {
	LevelBase
	Content []*tl.Text
}

func (l *titleLevel) Draw(screen *tl.Screen) {
	w, h := screen.Size()
	for row, t := range l.Content {
		t.SetPosition(utils.CenterText(t.Text(), w, h, row))
		t.Draw(screen)
	}
}

func (l *titleLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey && event.Key == tl.KeyEnter {
		l.ResultCh <- MainLevel(l.Context())
	}
}

func (l *titleLevel) Result() chan MinesweeperLevel {
	return l.ResultCh
}

func (l *titleLevel) Context() context.Context {
	return l.ctx
}

func TitleLevel(ctx context.Context) MinesweeperLevel {
	config, err := ConfigFromCtx(ctx)
	if err != nil {
		log.Print("Unable to decode config: ", err)
		return nil
	}

	fg, bg := config.FgColor, config.BgColor

	level := &titleLevel{
		Content: make([]*tl.Text, 0),
	}

	for _, str := range config.TitleContent {
		level.Content = append(level.Content, tl.NewText(0, 0, str, fg, bg))
	}

	level.Level = tl.NewBaseLevel(tl.Cell{
		Bg: bg,
		Fg: fg,
	})
	level.ResultCh = make(chan MinesweeperLevel)
	level.ctx = ctx

	return level
}
