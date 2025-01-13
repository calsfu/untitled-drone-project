package mission

type Point struct {
	X float64
	Y float64
	Z float64
}

// NewPoint creates a new point with the given coordinates
func NewPoint(x, y, z float64) *Point {
	return &Point{
		X: x,
		Y: y,
		Z: z,
	}
}
