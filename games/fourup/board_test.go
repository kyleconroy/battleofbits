package fourup

import (
	"testing"
)

func TestBoardFull(t *testing.T) {
	t.Parallel()
	fullBoard := Board{
		[7]Piece{2, 2, 2, 2, 2, 2, 2},
		[7]Piece{2, 2, 2, 2, 2, 2, 2},
		[7]Piece{1, 2, 2, 2, 2, 2, 2},
		[7]Piece{1, 2, 2, 2, 2, 2, 2},
		[7]Piece{1, 2, 2, 2, 2, 2, 2},
		[7]Piece{1, 2, 2, 2, 2, 2, 2},
	}
	if !fullBoard.Full() {
		t.Errorf("Full board should be marked full")
	}

	boardWithRoom := Board{
		[7]Piece{2, 0, 2, 2, 2, 2, 2},
		[7]Piece{2, 2, 2, 2, 2, 2, 2},
		[7]Piece{1, 2, 2, 2, 2, 2, 2},
		[7]Piece{1, 2, 2, 2, 2, 2, 2},
		[7]Piece{1, 2, 2, 2, 2, 2, 2},
		[7]Piece{1, 2, 2, 2, 2, 2, 2},
	}
	if boardWithRoom.Full() {
		t.Errorf("Board with room be marked not full")
	}
}

func TestGameOver(t *testing.T) {
	t.Parallel()

	winThirdVertical := Board{
		[7]Piece{0, 0, 0, 0, 0, 0, 1},
		[7]Piece{0, 0, 0, 0, 0, 0, 1},
		[7]Piece{0, 0, 0, 0, 0, 0, 1},
		[7]Piece{1, 0, 0, 0, 0, 0, 1},
		[7]Piece{1, 0, 0, 0, 0, 0, 2},
		[7]Piece{1, 0, 0, 0, 0, 0, 2},
	}
	if !winThirdVertical.Over() {
		t.Errorf("Game should be over if 4 vertical tiles " +
			"starting in top row, form a connect four")
	}

	winVertical := Board{
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
	}
	if !winVertical.Over() {
		t.Errorf("Game should be over if 4 vertical tiles are in a row")
	}

	winOtherVertical := Board{
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 1},
		[7]Piece{0, 0, 0, 0, 0, 0, 1},
		[7]Piece{0, 0, 0, 0, 0, 0, 1},
		[7]Piece{0, 0, 0, 0, 0, 0, 1},
	}
	if !winOtherVertical.Over() {
		t.Errorf("Game should be over if 4 other vertical tiles are in a row")
	}

	winHorizontal := Board{
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 1, 1, 1, 1, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
	}
	if !winHorizontal.Over() {
		t.Errorf("Game should be over if 4 horizontal tiles are in a row")
	}

	winDiagonal := Board{
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 1, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 1, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 1, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 1, 0},
	}
	if !winDiagonal.Over() {
		t.Errorf("Game should be over if 4 southeast diagonal tiles are in a row")
	}

	winSouthwestDiagonal := Board{
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 1, 0, 0},
		[7]Piece{0, 0, 0, 1, 0, 0, 0},
		[7]Piece{0, 0, 1, 0, 1, 0, 0},
		[7]Piece{0, 1, 0, 0, 0, 1, 0},
	}
	if !winSouthwestDiagonal.Over() {
		t.Errorf("Game should be over if 4 southwest diagonal tiles are in a row")
	}

	unfinishedGame := Board{
		[7]Piece{0, 2, 0, 2, 0, 0, 0},
		[7]Piece{0, 0, 2, 2, 1, 1, 1},
		[7]Piece{0, 0, 1, 1, 2, 2, 2},
		[7]Piece{0, 0, 2, 1, 2, 1, 0},
		[7]Piece{0, 0, 0, 2, 1, 1, 0},
		[7]Piece{0, 0, 1, 1, 2, 2, 2},
	}
	if unfinishedGame.Over() {
		t.Errorf("Game was marked over, but wasn't over")
	}
}

func TestApplyMoveToBoard(t *testing.T) {
	t.Parallel()
	emptyBoard := Board{
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
	}

	oneMoveBoard := Board{
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 0},
		[7]Piece{0, 0, 0, 0, 0, 0, 1},
	}

	emptyBoard.Drop(Red, 6)
	if emptyBoard != oneMoveBoard {
		t.Errorf("New board does not equal board with expected move")
	}

	columnFullBoard := Board{
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
		[7]Piece{1, 0, 0, 0, 0, 0, 0},
	}

	err := columnFullBoard.Drop(Red, 0)
	if err.Error() != "No room in column 0 for a move" {
		t.Errorf("Should have rejected move in column 0, did not, error was %s", err.Error())
	}

	err = columnFullBoard.Drop(Red, -22)
	if err == nil || err.Error() != "Move -22 is invalid" {
		t.Errorf("Should have rejected negative move, did not, error was %s", err.Error())
	}

	err = columnFullBoard.Drop(Red, 7)
	if err == nil || err.Error() != "Move 7 is invalid" {
		t.Errorf("Should have rejected positive move, did not, error was %s", err.Error())
	}
}
