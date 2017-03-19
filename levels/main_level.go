package levels

import (
	"context"
	"log"

	tl "github.com/JoelOtter/termloop"
	"github.com/ryanbaer/minesweeper/components"
)

type mainLevel struct {
	tl.Level
	Board    *components.Board
	ResultCh chan MinesweeperLevel
	ctx      context.Context
}

func (t *mainLevel) Context() context.Context {
	return t.ctx
}

func (t *mainLevel) Draw(s *tl.Screen) {
	t.Board.Draw(s)
}

func (t *mainLevel) Tick(event tl.Event) {
	// t.Board.Tick(event)
	// if event.Type == tl.EventKey && event.Key == tl.KeyEnter {
	// 	t.ResultCh <- nil
	// }
	if event.Type == tl.EventMouse && event.Key == tl.MouseRelease {
		exploded := t.Board.Click(event.MouseX, event.MouseY)
		if exploded {
		}

	}
}

func (m *mainLevel) Result() chan MinesweeperLevel {
	return m.ResultCh
}

func MainLevel(ctx context.Context) MinesweeperLevel {
	config, err := ConfigFromCtx(ctx)
	if err != nil {
		log.Print("Unable to decode config: ", err)
		return nil
	}

	ml := &mainLevel{
		Level: tl.NewBaseLevel(tl.Cell{
			Fg: tl.ColorBlack,
			Bg: tl.ColorWhite,
		}),
		ResultCh: make(chan MinesweeperLevel),
	}

	board, err := components.NewBoard(config.Width, config.Height, config.Mines)
	if err != nil {
		log.Print("Unable to decode config: ", err)
		return nil
	}

	ml.Board = board

	return ml
}
