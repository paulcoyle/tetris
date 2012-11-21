package tetris

import (
	"fmt"
)

var leastRotations = [4][4]int{
	{0, 1, 2, -1},
	{-1, 0, 1, 2},
	{-2, -1, 0, 1},
	{1, -2, -1, 0},
}

func FindLeastRotation(from, to int) int {
	return leastRotations[from][to]
}

// For the given board, clears any full lines adjusting the board accordingly.
// The value returned is the number of lines cleared in the operation.  This
// method modifies directly the board passed in.  If this is not desired, be
// sure to Copy() the board before passing it in.
func ClearFullLines(b *Board) int {
	lines := FindFullLines(b)
	if len(lines) == 0 {
		return 0
	}

	step := 1
	workRow := lines[len(lines)-step] // start on lowest cleared line
	for (workRow - step) >= 0 {
		copyRow := workRow - step
		nextLineIdx := len(lines) - step - 1

		if nextLineIdx < 0 || copyRow != lines[nextLineIdx] {
			// If the copy source row is not the next clear line, copy it to
			// the working row then move the working row up by one
			b.CopyRow(copyRow, workRow)
			workRow--
		} else {
			// Otherwise, increase the step and retry on a copyRow step above
			// the workRow
			step = step + 1
		}
	}

	// Finally, clear the remaining rows
	for workRow >= 0 {
		b.SetRow(workRow, false)
		workRow--
	}

	return len(lines)
}

// Locates the rows which have full lines and returns their row indices in
// asending order.
func FindFullLines(b *Board) []int {
	lines := make([]int, 0)

	for row := 0; row < b.Height(); row++ {
		full := true
		for col := 0; col < b.Width(); col++ {
			if set, _ := b.Block(row, col); !set {
				full = false
				break
			}
		}

		if full {
			lines = append(lines, row)
		}
	}

	return lines
}

// Returns the board, the final column, and an error
func PlaceInColumn(b *Board, s *Tetromino, row, col int) (*Board, int, error) {
	err := CheckPlacement(b, s, row, col)
	if err != nil {
		return nil, -1, err
	}

	board, frow, err := PlaceInLastRow(b, s, row, col)
	if err != nil {
		return nil, -1, err
	}

	return board, frow, nil
}

// Returns the board, the final row, and an error
func PlaceInLastRow(b *Board, s *Tetromino, row, col int) (*Board, int, error) {
	for i := row; i < b.Height(); i++ {
		e := CheckPlacement(b, s, i+1, col)
		if e != nil {
			b, _ := Place(b, s, i, col)
			return b, i, nil
		}
	}
	return nil, -1, fmt.Errorf("Could not place in row %d!", row)
}

func Place(b *Board, t *Tetromino, row, col int) (*Board, error) {
	board := b.Copy()

	err := CheckPlacement(board, t, row, col)
	if err != nil {
		return board, err
	}

	origin_row := 1
	origin_col := 2

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r := row - origin_row + i
			c := col - origin_col + j
			if r >= 0 && r < board.Height() && c >= 0 && c < board.Width() && t.Data()[i][j] {
				err = board.SetBlock(r, c, true)
				if err != nil {
					return board, fmt.Errorf("Can't place block (%d,%d)!\n", r, c)
				}
			}
		}
	}

	return board, nil
}

func CheckPlacement(b *Board, t *Tetromino, row, col int) error {
	left := t.Bounds().Left
	right := t.Bounds().Right
	bottom := t.Bounds().Bottom

	origin_row := 1
	origin_col := 2

	if (left-origin_col+col) < 0 || (right-origin_col+col) >= b.Width() ||
		(bottom-origin_row+row) >= b.Height() {
		return fmt.Errorf("Block placed at (%d,%d) would be out of bounds!", row, col)
	}

	// Now that we know it's not out of bounds, loop through 4x4 data
	// Convert (i,j) to global position (x,y) and check if the board has been blocked there
	// If it has, return error

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r := row - origin_row + i
			c := col - origin_col + j
			d, err := b.Block(r, c)
			if t.Data()[i][j] && d && err == nil {
				return fmt.Errorf("Block (%d,%d) already taken!", r, c)
			}
		}
	}

	return nil
}
