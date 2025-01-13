package mission

type Zone struct {
	// Name of the zone. Should encode the type of zone
	Name string
	// Points that define the zone
	Points []Point
}

func (z *Zone) AddPoint(p Point) {
	z.Points = append(z.Points, p)
}

func (z *Zone) RemovePoint(p Point) {
	for i, point := range z.Points {
		if point == p {
			z.Points = append(z.Points[:i], z.Points[i+1:]...)
			break
		}
	}
}
