package tetris

import (
	"fmt"
	"testing"
)

func TestFindLeastRotation(t *testing.T) {
	A := 0
	B := 3
	expected := -1
	actual := FindLeastRotation(A, B)
	
	if (expected != actual) {
		t.Errorf("Least rotation from %d->%d should be %d, was %d", A, B, expected, actual)
	}

	A = 1
	B = 3
	expected = 2
	actual = FindLeastRotation(A, B)

	if (expected != actual) {
		t.Errorf("Least rotation from %d->%d should be %d, was %d", A, B, expected, actual)
	}

	A = 0
	B = 2
	expected = 2
	actual = FindLeastRotation(A, B)

	if (expected != actual) {
		t.Errorf("Least rotation from %d->%d should be %d, was %d", A, B, expected, actual)
	}

	A = 2
	B = 1
	expected = -1
	actual = FindLeastRotation(A, B)

	if (expected != actual) {
		t.Errorf("Least rotation from %d->%d should be %d, was %d", A, B, expected, actual)
	}

	A = 3
	B = 0
	expected = 1
	actual = FindLeastRotation(A, B)

	if (expected != actual) {
		t.Errorf("Least rotation from %d->%d should be %d, was %d", A, B, expected, actual)
	}

	A = 3
	B = 1
	expected = -2
	actual = FindLeastRotation(A, B)

	if (expected != actual) {
		t.Errorf("Least rotation from %d->%d should be %d, was %d", A, B, expected, actual)
	}

	A = 2
	B = 2
	expected = 0
	actual = FindLeastRotation(A, B)

	if (expected != actual) {
		t.Errorf("Least rotation from %d->%d should be %d, was %d", A, B, expected, actual)
	}
}

func TestFindFullLinesWhenNoneExist(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|    #|",
		"|    #|",
		"|  # #|",
		"| ## #|",
		"|# ###|",
	})

	expected := []int{}
	result := FindFullLines(board)
	if !intArrSame(expected, result) {
		t.Error("Result should be an empty array")
	}
}

func TestFindFullLinesWhenOneExists(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|    #|",
		"|    #|",
		"|  # #|",
		"| ## #|",
		"|#####|",
	})

	expected := []int{4}
	result := FindFullLines(board)
	if !intArrSame(expected, result) {
		t.Error("Result should contain one element (4)")
	}
}

func TestFindFullLinesWhenManyExist(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|    #|",
		"|#####|",
		"|#####|",
		"|### #|",
		"|#####|",
	})

	expected := []int{1, 2, 4}
	result := FindFullLines(board)
	if !intArrSame(expected, result) {
		t.Error("Result should contain three elements (1,2,4)")
	}
}

func TestClearFullLinesWhenNoneExist(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|    #|",
		"|    #|",
		"|  # #|",
		"| ## #|",
		"|# ###|",
	})

	expected := board.Copy()

	numLines := ClearFullLines(board)
	if numLines != 0 {
		t.Error("Number of cleared lines should be 0")
	}
	if !board.Equal(expected) {
		t.Error("Modified board should be unchanged")
	}
}

func TestClearFullLinesWhenOneExists(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|    #|",
		"|    #|",
		"|  # #|",
		"| ## #|",
		"|#####|",
	})

	expected, _ := StringArrayToBoard([]string{
		"|     |",
		"|    #|",
		"|    #|",
		"|  # #|",
		"| ## #|",
	})

	numLines := ClearFullLines(board)
	if numLines != 1 {
		t.Error("Number of cleared lines should be 1")
	}
	if !board.Equal(expected) {
		t.Error("Modified board should reflect removed lines")
	}
}

func TestClearFullLinesWhenManyExist(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|    #|",
		"|#   #|",
		"|#####|",
		"|### #|",
		"|#####|",
	})

	expected, _ := StringArrayToBoard([]string{
		"|     |",
		"|     |",
		"|    #|",
		"|#   #|",
		"|### #|",
	})

	numLines := ClearFullLines(board)
	if numLines != 2 {
		t.Error("Number of cleared lines should be 2")
	}
	if !board.Equal(expected) {
		t.Error("Modified board should reflect removed lines")
		t.Error(board.String())
	}
}

func TestClearFullLinesWhenManyExistSuccsesively(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|    #|",
		"|#####|",
		"|#####|",
		"|### #|",
		"|#####|",
	})

	expected, _ := StringArrayToBoard([]string{
		"|     |",
		"|     |",
		"|     |",
		"|    #|",
		"|### #|",
	})

	numLines := ClearFullLines(board)
	if numLines != 3 {
		t.Error("Number of cleared lines should be 3")
	}
	if !board.Equal(expected) {
		t.Error("Modified board should reflect removed lines")
		t.Error(board.String())
	}
}

func TestClearFullLinesWhenAllFull(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|#####|",
		"|#####|",
		"|#####|",
		"|#####|",
		"|#####|",
	})

	expected, _ := StringArrayToBoard([]string{
		"|     |",
		"|     |",
		"|     |",
		"|     |",
		"|     |",
	})

	numLines := ClearFullLines(board)
	if numLines != 5 {
		t.Error("Number of cleared lines should be 5")
	}
	if !board.Equal(expected) {
		t.Error("Modified board should reflect removed lines")
		t.Error(board.String())
	}
}

func BenchmarkClearFullLines(b *testing.B) {
	b.StopTimer()
	board, _ := StringArrayToBoard([]string{
		"|####### ##|",
		"|##########|",
		"|####    ##|",
		"|##########|",
		"|### ######|",
		"|##########|",
		"|##########|",
		"|######## #|",
		"|#### #####|",
		"|##########|",
		"|##########|",
		"|##########|",
		"|#####   ##|",
		"|##########|",
		"|#### #####|",
		"|##########|",
		"|##########|",
		"|##########|",
		"|######  ##|",
		"|##########|",
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		clear := board.Copy()
		b.StartTimer()
		ClearFullLines(clear)
	}
}

func TestPlacesOnClearBoard(t *testing.T) {
	b, _ := NewBoard(10, 20)
	p, _ := NewTetromino("O", 0)
	_, err := Place(b, p, 0, 5)
	if err != nil {
		fmt.Print(err)
		t.Error("Shouldn't receive error on good placement!")
	}
}

func TestPlaceFailsOnBadPlacement(t *testing.T) {
	b, _ := NewBoard(10, 20)
	b.SetBlock(0, 3, true)
	p, _ := NewTetromino("I", 0)
	_, placement := Place(b, p, 0, 5)
	if placement == nil {
		t.Error("Shouldn't be able to do a bad placement!")
	}
}

func TestPlaceSucceedsOnValidPlacementI(t *testing.T) {
	b, _ := NewBoard(10, 20)
	b.SetBlock(4, 1, true)
	b.SetBlock(3, 1, true)
	b.SetBlock(3, 2, true)
	b.SetBlock(3, 3, true)
	b.SetBlock(3, 4, true)
	b.SetBlock(3, 5, true)
	b.SetBlock(3, 6, true)
	b.SetBlock(4, 6, true)
	b.SetBlock(5, 1, true)
	b.SetBlock(5, 2, true)
	b.SetBlock(5, 3, true)
	b.SetBlock(5, 4, true)
	b.SetBlock(5, 5, true)
	b.SetBlock(5, 6, true)
	p, _ := NewTetromino("I", 0)
	_, placement := Place(b, p, 4, 4)
	if placement != nil {
		t.Error("Shouldn't be able to do a bad placement!")
	}
}

func TestPlaceSucceedsOnValidPlacementT(t *testing.T) {
	b, _ := NewBoard(10, 20)
	b.SetBlock(1, 2, true)
	b.SetBlock(1, 5, true)

	b.SetBlock(2, 2, true)
	b.SetBlock(2, 5, true)

	b.SetBlock(3, 2, true)
	b.SetBlock(3, 3, true)
	b.SetBlock(3, 5, true)

	b.SetBlock(4, 3, true)
	b.SetBlock(4, 4, true)
	b.SetBlock(4, 5, true)

	p, _ := NewTetromino("T", 3)
	_, placement := Place(b, p, 2, 4)
	if placement != nil {
		t.Error("Shouldn't be able to do a bad placement!")
	}
}

func TestPlaceInColumnOnEmptyBoard(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|     |",
		"|     |",
		"|     |",
		"|     |",
		"|     |",
	})

	expected, _ := StringArrayToBoard([]string{
		"|     |",
		"|     |",
		"|     |",
		"| #   |",
		"|###  |",
	})

	tet, _ := NewTetromino("T", 2)
	board, _, err := PlaceInColumn(board, tet, 1, 1)
	if err != nil {
		t.Error("Placement should be valid")
		t.FailNow() // board will be nil in this case
	}

	if !board.Equal(expected) {
		t.Error("Resultant board should equal expected")
		fmt.Println(expected.String())
		fmt.Println(board.String())
	}
}

func TestPlaceInColumnOnExistingPiece(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|     |",
		"|     |",
		"|     |",
		"| #   |",
		"|###  |",
	})

	expected, _ := StringArrayToBoard([]string{
		"|     |",
		"|     |",
		"|   # |",
		"| ### |",
		"|#### |",
	})

	tet, _ := NewTetromino("T", 3)
	board, _, err := PlaceInColumn(board, tet, 1, 3)
	if err != nil {
		t.Error("Placement should be valid")
		t.FailNow() // board will be nil in this case
	}

	if !board.Equal(expected) {
		t.Error("Resultant board should equal expected")
		fmt.Println(expected.String())
		fmt.Println(board.String())
	}
}

func TestPlaceInColumnDoesntPhaseThroughBlocks(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|     |",
		"|     |",
		"|     |",
		"|     |",
		"|     |",
		"|    #|",
		"|     |",
		"|     |",
		"|     |",
		"|     |",
	})

	expected, _ := StringArrayToBoard([]string{
		"|     |",
		"|    #|",
		"|    #|",
		"|    #|",
		"|    #|",
		"|    #|",
		"|     |",
		"|     |",
		"|     |",
		"|     |",
	})

	tet, _ := NewTetromino("I", 1)
	board, _, err := PlaceInColumn(board, tet, 1, 4)
	if err != nil {
		t.Error("Placement should be valid")
		t.FailNow() // board will be nil in this case
	}

	if !board.Equal(expected) {
		t.Error("Resultant board should equal expected")
		fmt.Println(expected.String())
		fmt.Println(board.String())
	}
}

func TestPlaceInColumnErrorsOnBadPlacement(t *testing.T) {
	board, _ := StringArrayToBoard([]string{
		"|     |",
		"|  ## |",
		"|  ## |",
		"| ### |",
		"|#### |",
	})

	tet, _ := NewTetromino("T", 3)
	board, _, err := PlaceInColumn(board, tet, 1, 3)
	if err == nil {
		t.Error("Placement should be invalid")
	}
}

func intArrSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
