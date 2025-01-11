package robots

// Interface method implementation and attributes used by all robots
type BaseRobot struct {
	Id              string
	Position        [3]float64
	DescriptionText string
	Status          string
	BatteryLevel    float64
}

func (r *BaseRobot) GetId() string {
	return r.Id
}

func (r *BaseRobot) GetPosition() [3]float64 {
	return r.Position
}

func (r *BaseRobot) GetDescription() string {
	return r.DescriptionText
}

func (r *BaseRobot) GetStatus() string {
	return r.Status
}

func (r *BaseRobot) GetBatteryLevel() float64 {
	return r.BatteryLevel
}

// Only moving in the X and Y axis, since Z axis is only used for flying
func (r *BaseRobot) Move(x, y float64) error {
	r.Position[0] += x
	r.Position[1] += y
	return nil
}
