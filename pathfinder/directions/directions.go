package directions

// Direction is a direction in the grid (row, col)
type Direction = [2]int

func NewDirection(row, col int) Direction {
	return Direction{row, col}
}

var BOTTOM = NewDirection(1, 0)
var BOTTOM_LEFT = NewDirection(1, -1)
var BOTTOM_RIGHT = NewDirection(1, 1)

var LEFT = NewDirection(0, -1)
var RIGHT = NewDirection(0, 1)

var TOP = NewDirection(-1, 0)
var TOP_LEFT = NewDirection(-1, -1)
var TOP_RIGHT = NewDirection(-1, 1)

var DIRS_4 = []Direction{
	TOP,
	RIGHT,
	BOTTOM,
	LEFT,
}

var DIRS_8 = []Direction{
	TOP,
	RIGHT,
	BOTTOM,
	LEFT,
	TOP_RIGHT,
	BOTTOM_RIGHT,
	BOTTOM_LEFT,
	TOP_LEFT,
}
