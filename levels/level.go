package levels

import (
	"context"
	"errors"
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

type MinesweeperLevel interface {
	tl.Level
	Result() chan MinesweeperLevel
	Context() context.Context
}

type Config struct {
	Height       int
	Width        int
	Mines        int
	TitleContent []string
	WinContent   []string
	LoseContent  []string
	QuitMessage  string
}

const CtxConfig = "config"

var ErrConfig = errors.New("Unable to decode config")

func ConfigFromCtx(ctx context.Context) (*Config, error) {
	var (
		config *Config
		ok     bool
	)

	if config, ok = ctx.Value(CtxConfig).(*Config); !ok {
		return nil, ErrConfig
	}

	return config, nil
}

func StartGame(c *Config) {
	ctx := context.WithValue(context.Background(), CtxConfig, c)

	g := tl.NewGame()

	go func() {
		var cur = TitleLevel(ctx)

		for {
			g.Screen().SetLevel(cur)
			ch := cur.Result()
			next := <-ch
			if next == nil {
				fmt.Println("Received nil level, breaking")
				break
			}

			cur = next
		}
	}()

	g.Start()
}
