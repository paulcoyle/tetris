package tetris

import (
	"testing"
)

func TestConvertStringArrToBoardInconsistent(t *testing.T) {
	board, err := StringArrayToBoard([]string{
		"|# # # #|",
		"|# # # ##|", // <- width different from first row which sets the board width
	})
	if board != nil {
		t.Errorf("Returned board should be nil")
	}
	if err == nil {
		t.Errorf("An error should be returned")
	}
}

func TestConvertStringArrToBoard(t *testing.T) {
	board, err := StringArrayToBoard([]string{
		"|   |",
		"|###|",
		"|# #|",
	})
	if board == nil {
		t.Errorf("Returned board should not be nil")
	}
	if err != nil {
		t.Errorf("No error should be returned")
	}

	expected, _ := NewBoard(3, 3)
	expected.SetRow(1, true)
	expected.SetRow(2, true)
	expected.SetBlock(2, 1, false)
	if !board.Equal(expected) {
		t.Errorf("Converted board does not match expected")
	}
}
