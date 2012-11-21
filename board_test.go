package tetris

import (
	"testing"
)

func TestNewWidthErrorOnLessThanOne(t *testing.T) {
	board, err := NewBoard(0, 10)
	if err == nil {
		t.Error("An error should be produced when width is < 1")
	}
	if board != nil {
		t.Error("The board should be nil")
	}
}

func TestNewHeightErrorOnLessThanOne(t *testing.T) {
	board, err := NewBoard(10, 0)
	if err == nil {
		t.Error("An error should be produced when height is < 1")
	}
	if board != nil {
		t.Error("The board should be nil")
	}
}

func TestNewWithCorrectDimsIsOk(t *testing.T) {
	board, err := NewBoard(10, 10)
	if err != nil {
		t.Error("No error should be returned")
	}
	if board == nil {
		t.Error("The supplied board should not be nil")
	}
}

func TestErrorOnSetBlockOutsideWidthRange(t *testing.T) {
	board, _ := NewBoard(10, 10)
	err := board.SetBlock(0, -1, true)
	if err == nil {
		t.Error("An error should be returned when column below 0")
	}
	err = board.SetBlock(0, 10, true)
	if err == nil {
		t.Error("An error should be returned when column above width")
	}
}

func TestErrorOnSetBlockOutsideHeightRange(t *testing.T) {
	board, _ := NewBoard(10, 10)
	err := board.SetBlock(-1, 0, true)
	if err == nil {
		t.Error("An error should be returned when row below 0")
	}
	err = board.SetBlock(10, 0, true)
	if err == nil {
		t.Error("An error should be returned when row above height")
	}
}

func TestSetBlockWithinRange(t *testing.T) {
	board, _ := NewBoard(10, 10)
	board.SetBlock(0, 0, true)
	block, _ := board.Block(0, 0)
	if block != true {
		t.Error("Block data should be set correctly")
	}
}

func TestErrorOnGetBlockOutsideWidthRange(t *testing.T) {
	board, _ := NewBoard(10, 10)
	_, err := board.Block(0, -1)
	if err == nil {
		t.Error("An error should be returned when column below 0")
	}
	_, err = board.Block(0, 10)
	if err == nil {
		t.Error("An error should be returned when column above width")
	}
}

func TestErrorOnGetBlockOutsideHeightRange(t *testing.T) {
	board, _ := NewBoard(10, 10)
	_, err := board.Block(-1, 0)
	if err == nil {
		t.Error("An error should be returned when row below 0")
	}
	_, err = board.Block(10, 0)
	if err == nil {
		t.Error("An error should be returned when row above height")
	}
}

func TestGetBlockWithinRange(t *testing.T) {
	board, _ := NewBoard(10, 10)
	board.SetBlock(0, 0, true)
	block, _ := board.Block(0, 0)
	if block != true {
		t.Error("Block value returned should be true")
	}
	block, _ = board.Block(1, 1)
	if block != false {
		t.Error("Block value returned should be false")
	}
}

func TestSetRow(t *testing.T) {
	board, _ := NewBoard(5, 5)
	err := board.SetRow(4, true)
	if err != nil {
		t.Error("No error should be returned")
	}

	expected, _ := NewBoard(5, 5)
	expected.SetRow(4, true)

	if !board.Equal(expected) {
		t.Error("Board should equal expected")
	}
}

func TestGetDimensions(t *testing.T) {
	board, _ := NewBoard(3, 5)
	if board.Width() != 3 {
		t.Error("Board width should be 3")
	}
	if board.Height() != 5 {
		t.Error("Board height should be 5")
	}
}

func TestEquality(t *testing.T) {
	a, _ := NewBoard(3, 5)
	b, _ := NewBoard(3, 5)

	a.SetBlock(2, 0, true)
	a.SetBlock(2, 1, true)
	if a.Equal(b) {
		t.Errorf("Boards should not be equal")
	}

	b.SetBlock(2, 0, true)
	b.SetBlock(2, 1, true)
	if !a.Equal(b) {
		t.Errorf("Boards should be equal")
	}
}

func TestEqualityFailsOnDimensionMismatch(t *testing.T) {
	a, _ := NewBoard(3, 5)
	b, _ := NewBoard(3, 4)
	c, _ := NewBoard(4, 5)

	if a.Equal(b) {
		t.Errorf("Mismatch on height should be considered unequal")
	}

	if a.Equal(c) {
		t.Errorf("Mismatch on width should be considered unequal")
	}
}

func TestCopy(t *testing.T) {
	a, _ := NewBoard(5, 5)

	b := a.Copy()
	if !b.Equal(a) {
		t.Error("Copied board should be equal to it's source")
	}

	a.SetBlock(0, 0, true)
	if set, _ := b.Block(0, 0); set {
		t.Error("Changes to source should not affect the copy")
	}

	b.SetBlock(0, 1, true)
	if set, _ := a.Block(0, 1); set {
		t.Error("Changes to the copy should not affect the source")
	}
}
