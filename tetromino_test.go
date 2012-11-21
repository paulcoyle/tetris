package tetris

import (
	"testing"
)

func TestTetDataEqualityIdentifiesEqual(t *testing.T) {
	var a = &TetrominoData{
		{false, false, true, true},
		{false, false, true, true},
		{true, false, false, false},
		{false, false, false, false},
	}
	var b = &TetrominoData{
		{false, false, true, true},
		{false, false, true, true},
		{true, false, false, false},
		{false, false, false, false},
	}

	if !a.Equal(b) {
		t.Error("Data 'a' should be equal to 'b'")
	}

	if !b.Equal(a) {
		t.Error("Data 'b' should be equal to 'a' (reciprocal)")
	}
}

func TestTetDataEqualityIdentifiesInequal(t *testing.T) {
	var a = &TetrominoData{
		{false, false, true, true},
		{false, false, true, true},
		{true, false, false, false},
		{false, false, false, false},
	}
	var b = &TetrominoData{
		{false, false, false, true},
		{false, false, true, true},
		{true, false, false, false},
		{false, false, false, false},
	}

	if a.Equal(b) {
		t.Error("Data 'a' should be unequal to 'b'")
	}

	if b.Equal(a) {
		t.Error("Data 'b' should be unequal to 'a' (reciprocal)")
	}
}

func TestBuildTetDataFromInteger(t *testing.T) {
	var expected = &TetrominoData{
		{true, false, true, false},
		{false, true, false, true},
		{true, false, true, false},
		{false, true, false, true},
	}

	result := intToTetData(0xa5a5)

	if !result.Equal(expected) {
		t.Error("Result should equal expected")
	}
}

func TestBuildTetDataFromIntegerEvenWhenHuge(t *testing.T) {
	var expected = &TetrominoData{
		{true, false, true, false},
		{false, true, false, true},
		{true, false, true, false},
		{false, true, false, true},
	}

	result := intToTetData(0xffa5a5)

	if !result.Equal(expected) {
		t.Error("Result should equal expected")
	}
}

func TestNewTetInvalidKind(t *testing.T) {
	tet, err := NewTetromino("X", 0)
	if err == nil {
		t.Error("An error should be non-nil for an invalid tetromino kind")
	}
	if tet != nil {
		t.Error("Returned tetromino should be nil")
	}
}

func TestNewTetInvalidOrientation(t *testing.T) {
	tet, err := NewTetromino("I", 2)
	if err == nil {
		t.Error("An error should be non-nil for an invalid tetromino orientation")
	}
	if tet != nil {
		t.Error("Returned tetromino should be nil")
	}
}

func TestNewTetNegativeOrientation(t *testing.T) {
	tet, err := NewTetromino("I", -1)
	if err == nil {
		t.Error("An error should be non-nil for a negative tetromino orientation")
	}
	if tet != nil {
		t.Error("Returned tetromino should be nil")
	}
}

func TestNewTetHasCorrectData(t *testing.T) {
	tet, err := NewTetromino("I", 1)
	if err != nil {
		t.Error("An error should not be returned for a valid tetromino")
	}
	if tet == nil {
		t.Error("Returned tetromino should not be nil")
	}

	expected := intToTetData(tet_orients["I"][1])
	if !tet.data.Equal(expected) {
		t.Error("New tetromino data should equal expected")
	}
}

// Forward meaning incrementing, internally, the orientation by one
func TestRotateTetForward(t *testing.T) {
	tet, err := NewTetromino("L", 0)
	if err != nil {
		t.Error("An error should not be returned for a valid tetromino")
	}
	if tet == nil {
		t.Error("Returned tetromino should not be nil")
	}

	orient := tet.RotateFwd()
	if orient != 1 {
		t.Error("Orientation returned from rotation should be 1")
	}

	expected := intToTetData(tet_orients["L"][1])
	if !tet.data.Equal(expected) {
		t.Error("New tetromino data should equal expected")
	}
}

// Checks that "overflowing" (going over max orientations) wraps back around
func TestRotateTetForwardOverflow(t *testing.T) {
	tet, err := NewTetromino("L", 3)
	if err != nil {
		t.Error("An error should not be returned for a valid tetromino")
	}
	if tet == nil {
		t.Error("Returned tetromino should not be nil")
	}

	orient := tet.RotateFwd()
	if orient != 0 {
		t.Error("Orientation returned from rotation should be 0")
	}

	expected := intToTetData(tet_orients["L"][0])
	if !tet.data.Equal(expected) {
		t.Error("New tetromino data should equal expected")
	}
}

// Backward meaning decrementing, internally, the orientation by one
func TestRotateTetBackward(t *testing.T) {
	tet, err := NewTetromino("L", 1)
	if err != nil {
		t.Error("An error should not be returned for a valid tetromino")
	}
	if tet == nil {
		t.Error("Returned tetromino should not be nil")
	}

	orient := tet.RotateBack()
	if orient != 0 {
		t.Error("Orientation returned from rotation should be 0")
	}

	expected := intToTetData(tet_orients["L"][0])
	if !tet.data.Equal(expected) {
		t.Error("New tetromino data should equal expected")
	}
}

// Checks that "underflowing" (going less than orient 0) wraps back around to end
func TestRotateTetBackwardUnderflow(t *testing.T) {
	tet, err := NewTetromino("L", 0)
	if err != nil {
		t.Error("An error should not be returned for a valid tetromino")
	}
	if tet == nil {
		t.Error("Returned tetromino should not be nil")
	}

	orient := tet.RotateBack()
	if orient != 3 {
		t.Error("Orientation returned from rotation should be 3")
	}

	expected := intToTetData(tet_orients["L"][3])
	if !tet.data.Equal(expected) {
		t.Error("New tetromino data should equal expected")
	}
}

func TestTetBoundaries(t *testing.T) {
	tet, err := NewTetromino("L", 0)
	if err != nil {
		t.Error("An error should not be returned for a valid tetromino")
	}
	if tet == nil {
		t.Error("Returned tetromino should not be nil")
	}

	bounds := tet.Bounds()
	expected := &TetrominoBounds{1, 3, 1, 2}
	if !bounds.Equal(expected) {
		t.Error("Bounds do not match expected.")
	}

	tet.RotateFwd()
	bounds = tet.Bounds()
	expected = &TetrominoBounds{2, 3, 0, 2}
	if !bounds.Equal(expected) {
		t.Error("Bounds do not match expected after rotation.")
	}
}

func TestTetBoundaryEquality(t *testing.T) {
	a := &TetrominoBounds{1, 2, 3, 4}
	b := &TetrominoBounds{1, 2, 3, 4}

	if !a.Equal(b) {
		t.Error("Bounds should be equal")
	}
	if !b.Equal(a) {
		t.Error("Bounds should be equal reciprocally")
	}

	b.Top = 2

	if a.Equal(b) {
		t.Error("Bounds should be unequal")
	}
	if b.Equal(a) {
		t.Error("Bounds should be unequal reciprocally")
	}
}

func TestNumOrientsOnBadKind(t *testing.T) {
	num := NumTetOrients("X")
	if num != -1 {
		t.Error("Number of orients returned should be -1")
	}
}

func TestNumOrientsOnGoodKinds(t *testing.T) {
	num := NumTetOrients("O")
	if num != 1 {
		t.Error("Number of orients returned should be 1")
	}

	num = NumTetOrients("S")
	if num != 2 {
		t.Error("Number of orients returned should be 2")
	}

	num = NumTetOrients("J")
	if num != 4 {
		t.Error("Number of orients returned should be 4")
	}
}
