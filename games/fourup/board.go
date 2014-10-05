package fourup

import (
	"errors"
	"fmt"
)

type Piece int

const (
	NumRows        = 6
	NumColumns     = 7
	NumConsecutive = 4

	Empty Piece = 0
	Red   Piece = 1
	Black Piece = 2
)

type Board [NumRows][NumColumns]Piece

// row varies, column does not.
func (b Board) checkVerticalWin(column int) (bool, Piece) {
	checkRowInColumn := func(column int, row int) (bool, Piece) {
		initColor := b[row][column]
		for k := 0; k < NumConsecutive; k++ {
			if row+k >= NumRows {
				return false, initColor
			}
			value := b[row+k][column]
			if value == Empty || value != initColor {
				return false, value
			}
		}
		// if we get here and haven't broken, seen 4 in a row of the same color
		return true, initColor
	}

	for row := 0; row <= (NumRows - NumConsecutive); row++ {
		initColor := b[row][column]
		if initColor == Empty {
			continue
		}
		if over, winner := checkRowInColumn(column, row); over {
			return true, winner
		}
	}
	return false, Empty
}

func (b Board) checkHorizontalWin(row int) (bool, Piece) {
	checkColumnInRow := func(row int, column int) (bool, Piece) {
		initColor := b[row][column]
		for k := 0; k < NumConsecutive; k++ {
			if column+k >= NumColumns {
				return false, Empty
			}
			if b[row][column+k] != initColor {
				return false, Empty
			}
		}
		// if we get here and haven't broken, seen 4 in a row of the same color
		return true, initColor
	}
	for column := 0; column < NumConsecutive; column++ {
		initColor := b[row][column]
		if initColor == Empty {
			continue
		}
		if over, winner := checkColumnInRow(row, column); over {
			return true, winner
		}
	}
	return false, Empty
}

// check squares down and to the right for a match
func (b Board) checkSoutheastDiagonalWin(row, column int) (bool, Piece) {

	initColor := b[row][column]
	if initColor == Empty {
		return false, Empty
	}
	for i := 0; i < NumConsecutive; i++ {
		if b[row+i][column+i] != initColor {
			return false, Empty
		}
	}
	return true, initColor
}

func (b Board) checkSouthwestDiagonalWin(row, column int) (bool, Piece) {

	initColor := b[row][column]
	if initColor == Empty {
		return false, Empty
	}
	for i := 0; i < NumConsecutive; i++ {
		if b[row+i][column-i] != initColor {
			return false, Empty
		}
	}
	return true, initColor
}

// Checks if a connect four exists
// I'm sure there's some more efficient way to conduct these checks, but at
// modern computer speeds, it really doesn't matter
func (b Board) winner() (bool, Piece) {
	for column := 0; column < NumColumns; column++ {
		if over, winner := b.checkVerticalWin(column); over {
			return true, winner
		}
	}

	for row := 0; row < NumRows; row++ {
		if over, winner := b.checkHorizontalWin(row); over {
			return true, winner
		}
	}
	for row := 0; row <= (NumRows - NumConsecutive); row++ {
		for column := 0; column <= (NumColumns - NumConsecutive); column++ {
			if over, winner := b.checkSoutheastDiagonalWin(row, column); over {
				return true, winner
			}
		}
	}
	for column := NumColumns - NumConsecutive; column < NumColumns; column++ {
		for row := 0; row <= (NumRows - NumConsecutive); row++ {
			if over, winner := b.checkSouthwestDiagonalWin(row, column); over {
				return true, winner
			}
		}
	}
	return false, Empty
}

func (b Board) Over() bool {
	over, _ := b.winner()
	return over
}

func (b Board) Full() bool {
	// will check the top row, which is always the last to fill up.
	for column := 0; column < NumColumns; column++ {
		if b[0][column] == Empty {
			return false
		}
	}
	return true
}

// Returns error if the move is invalid
func (b *Board) Drop(color Piece, move int) error {
	if move >= NumColumns || move < 0 {
		return errors.New(fmt.Sprintf("Move %d is invalid", move))
	}
	for i := NumRows - 1; i >= 0; i-- {
		if b[i][move] == 0 {
			b[i][move] = color
			return nil
		}
	}
	return errors.New(fmt.Sprintf("No room in column %d for a move", move))
}
