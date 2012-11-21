package tetris

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	width  int
	height int
	data   [][]bool
}

func NewBoard(width, height int) (*Board, error) {
	if (width < 1) || (height < 1) {
		err := fmt.Errorf("Width and Height must both be greater than 1")
		return nil, err
	}

	board := &Board{width, height, make([][]bool, height)}
	for i := 0; i < height; i++ {
		board.data[i] = make([]bool, width)
	}

	return board, nil
}

func (b *Board) Width() int {
	return b.width
}

func (b *Board) Height() int {
	return b.height
}

// Used internally to check that the given row and column lie within the board's
// boundaries.
func (b *Board) checkBlockRange(row, col int) error {
	if (row < 0) || (col < 0) || (row >= b.height) || (col >= b.width) {
		err := fmt.Errorf("Block row:%d, col:%d is out of range", row, col)
		return err
	}

	return nil
}

// Returns the set value for a block (true if set, false if not).
func (b *Board) Block(row, col int) (bool, error) {
	err := b.checkBlockRange(row, col)
	if err != nil {
		return false, err
	}

	return b.data[row][col], nil
}

// Sets the given block to the supplied set value.
func (b *Board) SetBlock(row, col int, value bool) error {
	err := b.checkBlockRange(row, col)
	if err != nil {
		return err
	}

	b.data[row][col] = value

	return nil
}

// Sets an entire rwo to the given value.
func (b *Board) SetRow(row int, value bool) error {
	err := b.checkBlockRange(row, 0)
	if err != nil {
		return err
	}

	for col := 0; col < b.width; col++ {
		b.data[row][col] = value
	}

	return nil
}

// Sets an entire column to the given value.
func (b *Board) SetCol(col int, value bool) error {
	err := b.checkBlockRange(0, col)
	if err != nil {
		return err
	}

	for row := 0; row < b.height; row++ {
		b.data[row][col] = value
	}

	return nil
}

// Copies one row (from) to another (to).
func (b *Board) CopyRow(from, to int) {
	// TODO check bounds
	for col := 0; col < b.Width(); col++ {
		b.data[to][col] = b.data[from][col]
	}
}

// Creates a copy of the current board.
func (b *Board) Copy() *Board {
	c, _ := NewBoard(b.width, b.height)
	for row := 0; row < b.height; row++ {
		for col := 0; col < b.width; col++ {
			c.data[row][col] = b.data[row][col]
		}
	}
	return c
}

// Determines equality between two boards.
func (b *Board) Equal(other *Board) bool {
	if b.width != other.width || b.height != other.height {
		return false
	}

	for row := 0; row < b.height; row++ {
		for col := 0; col < b.width; col++ {
			if b.data[row][col] != other.data[row][col] {
				return false
			}
		}
	}
	return true
}

// Outputs the board as a string that kinda sorta looks like a tetris board.
func (b *Board) String() string {
	var out string = "  ┌"
	out += strings.Repeat("─", b.width)
	out += "┐\n"
	for row := 0; row < b.height; row++ {
		r := strconv.Itoa(row)
		out += r
		if row < 10 {
			out += " "
		}

		out += "│"
		for col := 0; col < b.width; col++ {
			if b.data[row][col] {
				out += "X"
			} else {
				out += " "
			}
		}
		out += "│\n"
	}
	out += "  └"
	out += strings.Repeat("─", b.width)
	out += "┘\n"

	out += "   "
	for i := 0; i < b.width; i++ {
		r := strconv.Itoa(i)
		out += r
	}
	out += "\n"

	return out
}
