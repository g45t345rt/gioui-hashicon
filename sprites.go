package gioui_hashicon

type Light string

const (
	Top   Light = "top"
	Left  Light = "left"
	Right Light = "right"
)

type Sprite struct {
	x      float32
	y      float32
	shape  uint
	hidden bool
	light  Light
}

var Sprites = []Sprite{
	{x: 0, y: 0, shape: 1, hidden: true},
	{x: 0, y: 0, shape: 0, light: Top},
	{x: 0, y: 0.25, shape: 1, light: Left},
	{x: 0, y: 0.25, shape: 0, light: Left},
	{x: 0, y: 0.5, shape: 1, light: Left},
	{x: 0, y: 0.5, shape: 0, light: Left},
	{x: 0, y: 0.75, shape: 1, hidden: true},

	{x: 0.25, y: -0.125, shape: 0, light: Top},
	{x: 0.25, y: 0.125, shape: 1, light: Top},
	{x: 0.25, y: 0.125, shape: 0, light: Top},
	{x: 0.25, y: 0.375, shape: 1, light: Left},
	{x: 0.25, y: 0.375, shape: 0, light: Left},
	{x: 0.25, y: 0.625, shape: 1, light: Left},
	{x: 0.25, y: 0.625, shape: 0, light: Left},

	{x: 0.5, y: 0, shape: 1, light: Top},
	{x: 0.5, y: 0, shape: 0, light: Top},
	{x: 0.5, y: 0.25, shape: 1, light: Top},
	{x: 0.5, y: 0.25, shape: 0, light: Right},
	{x: 0.5, y: 0.5, shape: 1, light: Right},
	{x: 0.5, y: 0.5, shape: 0, light: Right},
	{x: 0.5, y: 0.75, shape: 1, light: Right},

	{x: 0.75, y: -0.125, shape: 0, hidden: true},
	{x: 0.75, y: 0.125, shape: 1, light: Top},
	{x: 0.75, y: 0.125, shape: 0, light: Right},
	{x: 0.75, y: 0.375, shape: 1, light: Right},
	{x: 0.75, y: 0.375, shape: 0, light: Right},
	{x: 0.75, y: 0.625, shape: 1, light: Right},
	{x: 0.75, y: 0.625, shape: 0, hidden: true},
}
