package tetris

import (
	"fmt"
)

// Contains the possible tetromino kinds and their respective orientations as
// given in the Pason guide.  Orientations are stored in integerial format and
// can be converted to TetrominoData types with intToTetData.
var tet_orients = map[string][]int{
	"O": {0x0660},
	"I": {0x0f00, 0x2222},
	"S": {0x0360, 0x2310},
	"Z": {0x0630, 0x1320},
	"L": {0x0740, 0x2230, 0x1700, 0x6220},
	"J": {0x0710, 0x3220, 0x4700, 0x2260},
	"T": {0x0720, 0x2320, 0x2700, 0x2620},
}

type TetrominoData [4][4]bool

func (d *TetrominoData) Equal(other *TetrominoData) bool {
	equal := true

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if d[row][col] != other[row][col] {
				equal = false
				break
			}
		}
		if !equal {
			break
		}
	}

	return equal
}

type TetrominoBounds struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

// Returns the number of possible orientations for the given kind of tetromino.
// Will return -1 if the kind of tetromino does not exist.
func NumTetOrients(kind string) int {
	arr, ok := tet_orients[kind]

	if !ok {
		return -1
	}

	return len(arr)
}

// Determines equality between two TetrominoBounds
func (b *TetrominoBounds) Equal(other *TetrominoBounds) bool {
	if b.Top != other.Top || b.Bottom != other.Bottom || b.Left != other.Left || b.Right != other.Right {
		return false
	}
	return true
}

type Tetromino struct {
	kind   string
	orient int
	data   *TetrominoData
	bounds *TetrominoBounds
}

func NewTetromino(kind string, orient int) (*Tetromino, error) {
	orients, ok := tet_orients[kind]
	if !ok {
		return nil, fmt.Errorf("Supplied tetromino kind %s is not valid", kind)
	}
	if orient < 0 || orient > (len(orients)-1) {
		return nil, fmt.Errorf("Orientation %d for tetromino kind %s is not valid", orient, kind)
	}

	tet := &Tetromino{kind, orient, nil, nil}
	tet.setData(intToTetData(orients[orient]))

	return tet, nil
}

func (t *Tetromino) Copy() *Tetromino {
	cpy := new(Tetromino)
	*cpy = *t
	return cpy
}

func (t *Tetromino) Kind() string {
	return t.kind
}

func (t *Tetromino) Orient() int {
	return t.orient
}

// Rotates "forward" (increments orientation by 1) and returns new orientation
func (t *Tetromino) RotateFwd() int {
	return t.rotate(1)
}

// Rotates "backward" (decrements orientation by 1) and returns new orientation
func (t *Tetromino) RotateBack() int {
	return t.rotate(-1)
}

func (t *Tetromino) rotate(delta int) int {
	olen := len(tet_orients[t.kind])
	t.orient = (t.orient + delta) % olen
	if t.orient < 0 {
		t.orient = t.orient + olen
	}
	t.setData(intToTetData(tet_orients[t.kind][t.orient]))
	return t.orient
}

func (t *Tetromino) Bounds() *TetrominoBounds {
	return t.bounds
}

func (t *Tetromino) Data() *TetrominoData {
	return t.data
}

func (t *Tetromino) setData(data *TetrominoData) {
	t.data = data
	t.updateBounds()
}

// For the current orientation, this method updates the bounds field of the
// Tetromino struct representing the bounds of the orientation.
func (t *Tetromino) updateBounds() {
	left, right, top, bot := 3, 0, 3, 0

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if t.data[row][col] {
				if col < left {
					left = col
				}
				if col > right {
					right = col
				}
				if row < top {
					top = row
				}
				if row > bot {
					bot = row
				}
			}
		}
	}

	t.bounds = &TetrominoBounds{left, right, top, bot}
}

func (t *Tetromino) String() string {
	out := ""
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if t.data[row][col] {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}

	return out
}

// Converts the integer representation of a tetromino into a TetrominoData
// array.
func intToTetData(orient int) *TetrominoData {
	var out TetrominoData
	var check int = 0x8000
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if (orient & check) > 0 {
				out[row][col] = true
			}
			check = check >> 1
		}
	}
	return &out
}
