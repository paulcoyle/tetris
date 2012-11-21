package tetris

import (
	"fmt"
	"math/rand"
	"time"
)

// Usage:
//
//     board, err := StringArrayToBoard([]string{
//        "|    #|",
//        "|    #|",
//        "|  # #|",
//        "| ## #|",
//        "|#####|",
//     })
//
func StringArrayToBoard(rows []string) (*Board, error) {
	width, height := len(rows[0])-2, len(rows)
	board, _ := NewBoard(width, height)

	for row := 0; row < height; row++ {
		if len(rows[row])-2 != width {
			return nil, fmt.Errorf("Bad board format, widths are not consistent")
		}

		values := rows[row][1 : width+1]
		for col := 0; col < width; col++ {
			if values[col] == '#' {
				board.SetBlock(row, col, true)
			}
		}
	}

	return board, nil
}

// Generates a random tetromino and returns a reference to it.
func RandomTetromino() *Tetromino {
	rand.Seed(time.Now().UnixNano())
	keys := make([]string, 0)
	for k, _ := range tet_orients {
		keys = append(keys, k)
	}
	kind := keys[rand.Intn(len(keys))]
	orient := rand.Intn(NumTetOrients(kind))
	tet, _ := NewTetromino(kind, orient)
	return tet
}
