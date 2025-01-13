package mission

type RectZone struct {
	Zone    // Embedding the Zone struct
	TopLeft Point
	Width   float64
	Height  float64
}

func (z *RectZone) Contains(p Point) bool {
	return p.X >= z.TopLeft.X && p.X <= z.TopLeft.X+z.Width &&
		p.Y >= z.TopLeft.Y && p.Y <= z.TopLeft.Y+z.Height
}

func (z *RectZone) Resize(width, height float64) {
	z.Width = width
	z.Height = height
}

// NewRectZone creates a new rectangular zone with the given coordinates
func NewRectZone(name string, topLeft Point, width, height float64) *RectZone {
	return &RectZone{
		Zone:    Zone{Name: name},
		TopLeft: topLeft,
		Width:   width,
		Height:  height,
	}
}
