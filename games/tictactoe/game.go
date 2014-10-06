package tictactoe

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/kyleconroy/battleofbits/games"
)

type Piece int8
type Board [9]Piece

const (
	X Piece = 1
	O Piece = 2
)

func (p Piece) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	}
	return "-"
}

func (b Board) String() string {
	repr := `
    %s%s%s
    %s%s%s
    %s%s%s
    `
	return fmt.Sprintf(repr, b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8])
}

type Payload struct {
	Piece Piece `json:"piece"`
	Board Board `json:"board"`
}

type Move struct {
	Space int `json:"space"`
}

// A tic-tac-toe board is represented as nine-length array of integers
type Player struct {
	Piece Piece
	URL   *url.URL
}

func (p Player) NextMove(pc Piece, board Board) (int, error) {
	var move Move
	err := games.FetchMove(p.URL, Payload{Piece: pc, Board: board}, &move)
	if err != nil {
		return 0, err
	}
	return move.Space, nil
}

// Match Handlers
func random(p Piece, board Board) (int, error) {
	for i := 0; i < 20; i++ {
		if i := rand.Intn(9); board[i] == 0 {
			return i, nil
		}
	}
	return 0, fmt.Errorf("tried 20 random spaces")
}

func greedy(p Piece, board Board) (int, error) {
	for i, slot := range board {
		if slot == 0 {
			return i, nil
		}
	}
	return 0, fmt.Errorf("no open spaces")
}

type Match struct {
	Board     Board
	PlayerOne Player
	PlayerTwo Player
	Current   Piece
}

func (m *Match) Add(piece Piece, i int) error {
	if i < 0 || i > 8 {
		return fmt.Errorf("space %d is out of bounds", i)
	}
	if m.Board[i] != 0 {
		return fmt.Errorf("space %s is occupied by %d", i, m.Board[i])
	}
	m.Board[i] = piece
	return nil
}

// Maybe there should be a tick function that returns the board json and move json?
func (m *Match) Tick() error {
	var current Player
	if m.Current == X {
		current = m.PlayerOne
	} else {
		current = m.PlayerTwo
	}
	square, err := current.NextMove(current.Piece, m.Board)
	if err != nil {
		return err
	}

	err = m.Add(current.Piece, square)
	if err != nil {
		return err
	}

	// Toggle the active player
	if m.Current == X {
		m.Current = O
	} else {
		m.Current = X
	}
	return nil
}

func (m Match) Over() bool {
	match := func(i, j, k int) bool {
		return m.Board[i] == m.Board[j] &&
			m.Board[j] == m.Board[k] &&
			m.Board[i] != 0
	}
	wins := []bool{
		match(0, 1, 2), // horizontal rows
		match(3, 4, 5),
		match(6, 7, 8),
		match(0, 3, 6), // vertical rows
		match(1, 4, 7),
		match(2, 5, 8),
		match(0, 4, 8), // diagonal rows
		match(2, 4, 6),
	}
	for _, win := range wins {
		if win {
			return true
		}
	}
	for _, piece := range m.Board {
		if piece == 0 {
			return false
		}
	}
	return false
}

func NewMatch(one, two *url.URL) Match {
	return Match{
		Board:     Board{},
		PlayerOne: Player{Piece: X, URL: one},
		PlayerTwo: Player{Piece: O, URL: two},
		Current:   X,
	}
}
