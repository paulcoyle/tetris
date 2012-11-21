package tetris

import (
	"testing"
)

func TestConvertStringToBoardEmpty(t *testing.T) {
	expected, _ := NewBoard(10, 20)

	data := "00000000000000000000000000000000000000000000000000"
	board, err := StringToBoard(data)
	if err != nil {
		t.Errorf("No error should be returned")
	}
	if board == nil {
		t.Errorf("A non-nil board should be returned")
	}

	if !board.Equal(expected) {
		t.Errorf("Converted board does not match expected layout")
	}
}

func TestConvertStringToBoard(t *testing.T) {
	expected, _ := NewBoard(10, 20)
	expected.SetRow(19, true)

	data := "000000000000000000000000000000000000000000000003ff"
	board, err := StringToBoard(data)
	if err != nil {
		t.Errorf("No error should be returned")
	}
	if board == nil {
		t.Errorf("A non-nil board should be returned")
	}

	if !board.Equal(expected) {
		t.Errorf("Converted board does not match expected layout")
	}
}

func TestConvertStringToBoardTooLittleData(t *testing.T) {
	data := "000f"
	board, err := StringToBoard(data)
	if board != nil {
		t.Errorf("Returned board should be nil")
	}
	if err == nil {
		t.Errorf("An error should be returned")
	}
}

func TestConvertStringToBoardTooMuchData(t *testing.T) {
	data := "00000000000000000000000000000000000000000000000000f"
	board, err := StringToBoard(data)
	if board != nil {
		t.Errorf("Returned board should be nil")
	}
	if err == nil {
		t.Errorf("An error should be returned")
	}
}

func TestConvertStringToBoardBadData(t *testing.T) {
	data := "M0000000000000000000000000000000000000000000000000"
	board, err := StringToBoard(data)
	if board != nil {
		t.Errorf("Returned board should be nil")
	}
	if err == nil {
		t.Errorf("An error should be returned")
	}
}

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
