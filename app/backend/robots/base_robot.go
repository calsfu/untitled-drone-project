package robots

import (
	"drone-rescue/utils"
	"gonum.org/v1/gonum/mat"
)

// Interface method implementation and attributes used by all robots
type BaseRobot struct {
	Id              string
	DesiredPosition mat.VecDense // Desired position of the robot
	ActualPosition  mat.VecDense // Actual position of the robot
	DescriptionText string
	Status          string
	BatteryLevel    float64
}

func (r *BaseRobot) GetId() string {
	return r.Id
}

func (r *BaseRobot) GetDesiredPosition() mat.VecDense {
	return r.DesiredPosition
}

func (r *BaseRobot) GetActualPosition() mat.VecDense {
	return r.ActualPosition
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
func (r *BaseRobot) Move(dX, dY float64) float64 {
	r.DesiredPosition.SetVec(utils.X_INDEX, r.DesiredPosition.At(utils.X_INDEX, 0)+dX)
	r.DesiredPosition.SetVec(utils.Y_INDEX, r.DesiredPosition.At(utils.Y_INDEX, 0)+dY)

	return r.PositionError()
}

// Find the absolute value between the desired position and the actual position
func (r *BaseRobot) PositionError() float64 {
	error := mat.NewVecDense(3, nil)
	error.SubVec(&r.DesiredPosition, &r.ActualPosition)
	norm := mat.Norm(error, 2)
	return norm
}
