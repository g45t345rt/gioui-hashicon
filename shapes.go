package gioui_hashicon

type Triangle struct {
	x1 float32
	y1 float32
	x2 float32
	y2 float32
	x3 float32
	y3 float32
}

var Shapes = []Triangle{
	{x1: 0, y1: 0.25, x2: 0.25, y2: 0.125, x3: 0.25, y3: 0.375},
	{x1: 0, y1: 0, x2: 0.25, y2: 0.125, x3: 0, y3: 0.25},
}
