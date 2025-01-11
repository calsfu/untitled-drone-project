package robots

// Commands specific to drones, mainly using Mavlink
type FlyingRobot struct {
	BaseRobot // Embedding BaseRobot
	// Additional fields specific to FlyingRobot
	AltitudeVelocity float64
	YawVelocity      float64
}

func (r *FlyingRobot) Ascend() error {
	r.Status = "Ascending"
	r.Position[2] += 1 // Increasing hanging the Z position
	return nil
}

func (r *FlyingRobot) Descend() error {
	r.Status = "Descending"
	r.Position[2] -= 1 // Decreasing hanging the Z position
	return nil
}

func (r *FlyingRobot) Move(x, y, z float64) error {
	r.Position[0] += x
	r.Position[1] += y
	r.Position[2] += z
	return nil
}
