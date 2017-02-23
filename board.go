package main

import "fmt"

const hor = "----"
const ver = "|"

func line(numCells int) string {
	if numCells == 0 {
		return ""
	}
	str := ver

	for i := 0; i < numCells; i++ {
		str += hor + ver
	}

	return str
}

func cell(value int) string {
	if value == 0 {
		return "    "
	}
	return fmt.Sprintf("%4d", value)
}

type Board4x4 struct {
	cells [16]int
}

func (board *Board4x4) String() interface{} {
	str := line(4) + "\n"
	empty := ver + cell(0) + ver + cell(0) + ver + cell(0) + ver + cell(0) + ver + "\n"

	for i := 0; i < 4; i++ {
		str += empty + line(4) + "\n"
	}
	return str
}
