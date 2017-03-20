package levels

import (
	"context"
	"log"

	tl "github.com/JoelOtter/termloop"
)

type loseLevel struct {
	LevelBase
	Content []*tl.Text
}

func (l *loseLevel) Draw(screen *tl.Screen) {
	w, h := screen.Size()
	for i, t := range l.Content {
		t.SetPosition(w/2-len(t.Text())/2, h/3+i)
		t.Draw(screen)
	}
}

func (l *loseLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey && event.Key == tl.KeyEnter {
		l.ResultCh <- MainLevel(l.Context())
	}
}

func (l *loseLevel) Result() chan MinesweeperLevel {
	return l.ResultCh
}

func (l *loseLevel) Context() context.Context {
	return l.ctx
}

func LoseLevel(ctx context.Context) MinesweeperLevel {
	config, err := ConfigFromCtx(ctx)
	if err != nil {
		log.Print("Unable to decode config: ", err)
		return nil
	}

	fg, bg := config.FgColor, config.BgColor

	level := &loseLevel{
		Content: make([]*tl.Text, 0),
	}

	for _, str := range config.LoseContent {
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
