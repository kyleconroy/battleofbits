package fourup

import (
	"errors"
	"net/url"

	"github.com/kyleconroy/battleofbits/games"
)

type Payload struct {
	Piece Piece `json:"piece"`
	Board Board `json:"board"`
}

type Move struct {
	Column int `json:"column"`
}

// A tic-tac-toe board is represented as nine-length array of integers
type Player struct {
	Piece Piece
	URL   *url.URL `json:"-"`
}

type Match struct {
	Board     Board
	PlayerOne Player
	PlayerTwo Player
	Current   Piece
}

func (m *Match) Tick() error {
	var current Player
	var move Move

	if m.Current == Red {
		current = m.PlayerOne
	} else {
		current = m.PlayerTwo
	}

	payload := Payload{Piece: current.Piece, Board: m.Board}
	err := games.FetchMove(current.URL, payload, &move)
	if err != nil {
		return err
	}

	if m.Over() {
		return errors.New("move not allowed as the board is full or the game is over")
	}

	err = m.Board.Drop(current.Piece, move.Column)
	if err != nil {
		return err
	}

	// Toggle the active player
	if m.Current == Red {
		m.Current = Black
	} else {
		m.Current = Black
	}
	return nil
}

func (m Match) Over() bool {
	return m.Board.Over() || m.Board.Full()
}

func NewMatch(one, two *url.URL) Match {
	return Match{
		Board:     Board{},
		PlayerOne: Player{Piece: Red, URL: one},
		PlayerTwo: Player{Piece: Black, URL: two},
		Current:   Red,
	}
}
