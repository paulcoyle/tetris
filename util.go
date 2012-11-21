package tetris

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Based on the Pason spec for board data.  All boards are 10x20.
//   - each char is a hex value whose bits represent which blocks are occupied
//   - blocks are listed for each column of each row starting from the top left
var boardStringLen int = 10 * 20 / 4

func StringToBoard(data string) (*Board, error) {
	if len(data) != boardStringLen {
		return nil, fmt.Errorf("String data length incorrect (expected %d)", boardStringLen)
	}

	board, _ := NewBoard(10, 20)
	reader := strings.NewReader(data)
	var blockNum int = 0
	for {
		c, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("Error reading byte from board data")
		}

		i, err := strconv.ParseInt(string(c), 16, 8)
		if err != nil {
			return nil, fmt.Errorf("Error parsing board data")
		}

		var mask int64 = 0x8
		for mask > 0 {
			var row int = blockNum / 10
			var col int = blockNum % 10
			board.SetBlock(row, col, (i&mask > 0))
			mask = mask >> 1
			blockNum++
		}
	}

	return board, nil
}

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
