package levels

import (
	"context"
	"errors"

	tl "github.com/JoelOtter/termloop"
	"github.com/ryanbaer/minesweeper-go/lib"
)

const QuitMessage = "Press [ctrl +c] to quit at anytime"

type Config struct {
	Height       int
	Width        int
	Mines        int
	TitleContent []string
	WinContent   []string
	LoseContent  []string
	FgColor      tl.Attr
	BgColor      tl.Attr
}

const CtxConfig = "config"

var ErrConfig = errors.New("Unable to decode config")

func ValidateConfig(c *Config) error {
	if c.Width < lib.MinWidth || c.Height < lib.MinHeight {
		return lib.ErrInvalidDimensions
	}

	maxMines := c.Width*c.Height - 1
	if c.Mines > maxMines {
		return lib.ErrTooManyMines
	}

	return nil
}

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
