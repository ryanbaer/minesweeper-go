package levels

import (
	"context"
	"fmt"
	"log"

	tl "github.com/JoelOtter/termloop"
	"github.com/ryanbaer/minesweeper-go/components"
	"github.com/ryanbaer/minesweeper-go/utils"
)

type mainLevel struct {
	LevelBase
	Board  *components.Board
	state  components.GameState
	Footer *components.Footer
}

// stateMap maps a `components.GameState` to a `LevelFunc`
// This provides a data-driven approach to routing to the next level based on the
// condition of the game
var stateMap = map[components.GameState]LevelFunc{
	components.GameStateLost: LoseLevel,
	components.GameStateWon:  WinLevel,
}

func (t *mainLevel) Context() context.Context {
	return t.ctx
}

func (t *mainLevel) Draw(s *tl.Screen) {
	t.Board.Draw(s)
	t.Footer.Draw(s)
}

func (t *mainLevel) SetState(state components.GameState) {
	t.state = state
	if !t.GameActive() {
		if t.GameWon() {
			t.Footer.SetContent([]string{
				"All mines cleared!",
				"Press [enter] to continue",
				"",
				fmt.Sprintf("Mines Remaining: %d", t.Board.Remaining()),
			})
		} else {
			t.Footer.SetContent([]string{
				"You exploded!",
				"Examine your mistake and press [enter] to continue",
				"",
				fmt.Sprintf("Mines Remaining: %d", t.Board.Remaining()),
			})
			t.Board.ToggleSolution()
		}
	}

}

func (t *mainLevel) Tick(event tl.Event) {
	if utils.MouseUp(event) {
		if t.state == components.GameStateActive {
			result := t.Board.Click(event.MouseX, event.MouseY)
			t.Footer.UpdateItem(0, fmt.Sprintf("Mines Remaining: %d", t.Board.Remaining()))
			t.SetState(result)
		}
	} else if utils.EnterPress(event) {
		if !t.GameActive() {
			next, ok := stateMap[t.state]
			if !ok {
				t.Footer.Append(fmt.Sprintf("Error: Unable to find next state for state: %d", t.state))
				return
			}
			t.ResultCh <- next(t.Context())
		}

	}

}

func (l *mainLevel) GameLost() bool {
	return l.state == components.GameStateLost
}

func (l *mainLevel) GameWon() bool {
	return l.state == components.GameStateWon
}

func (l *mainLevel) GameActive() bool {
	return l.state == components.GameStateActive
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

	fg, bg := config.FgColor, config.BgColor

	level := &mainLevel{
		state: components.GameStateActive,
	}
	level.Level = tl.NewBaseLevel(tl.Cell{
		Fg: fg,
		Bg: bg,
	})
	level.ResultCh = make(chan MinesweeperLevel)
	level.ctx = ctx

	board, err := components.NewBoard(config.Width, config.Height, config.Mines)
	if err != nil {
		log.Print("Unable to decode config: ", err)
		return nil
	}

	level.Board = board

	foot := []string{
		fmt.Sprintf("Mines Remaining: %d", board.Remaining()),
		"",
		"Click on a point in the grid to begin",
		QuitMessage,
	}

	level.Footer = components.NewFooter(foot, fg, bg, 5)

	return level
}
