package games

type Game string

const (
	TicTacToe Game = "tictactoe"
	FourUp    Game = "fourup"
)

type Match interface {
	Tick() error
	Over() bool
}

//TODO Make this better
func Play(match Match) error {
	for {
		err := match.Tick()
		if err != nil {
			return err
		}
		if match.Over() {
			return nil
		}
	}
}
