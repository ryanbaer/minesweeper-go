package levels

import (
	"context"
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

func StartGame(c *Config) error {
	if err := ValidateConfig(c); err != nil {
		return err
	}

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
	return nil
}
