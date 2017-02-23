package main

import (
	"testing"
	"fmt"
)

func assertExp(t *testing.T, noun string, expected string, result string) {
	if result != expected {
		t.Errorf("Expected %s:\n\n%s\n\nbut got %s:\n\n%s",
			noun,
			expected,
			noun,
			result,
		)
	}
}

func TestLine(t *testing.T) {
	assertExp(t, "line",
		"|----|----|----|----|",
		line(4),
	)
}

func TestCellZero(t *testing.T) {
	assertExp(t, "cell",
		"    ",
		cell(0),
	)
}

func TestCellNumbers(t *testing.T) {
	for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		assertExp(t, "cell",
			fmt.Sprintf("%4d", num),
			cell(num),
		)
	}
}

func TestBlankBoardOutput(t *testing.T) {
	board := new(Board4x4)
	expected := `|----|----|----|----|
|    |    |    |    |
|----|----|----|----|
|    |    |    |    |
|----|----|----|----|
|    |    |    |    |
|----|----|----|----|
|    |    |    |    |
|----|----|----|----|
`

	if result := board.String(); result != expected {
		t.Errorf("Expected board:\n\n%s\n\nbut got board:\n\n%s", expected, result)
	}
}
