package levels

import (
	"context"

	tl "github.com/JoelOtter/termloop"
)

const backgroundChar = ' '

type LevelBase struct {
	tl.Level
	ResultCh chan MinesweeperLevel
	ctx      context.Context
}

type MinesweeperLevel interface {
	tl.Level
	Result() chan MinesweeperLevel
	Context() context.Context
}

type LevelFunc func(context.Context) MinesweeperLevel
