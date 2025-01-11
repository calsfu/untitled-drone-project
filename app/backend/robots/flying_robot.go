package robots

import (
	"drone-rescue/utils"
)

// Commands specific to drones, mainly using Mavlink
type FlyingRobot struct {
	BaseRobot // Embedding BaseRobot
	// Additional fields specific to FlyingRobot
	AltitudeVelocity float64
	YawVelocity      float64
}

func (r *FlyingRobot) ChangeAltitude(dist float64) float64 {
	if dist < 0 {
		r.Status = "Descending"
	} else {
		r.Status = "Ascending"
	}

	r.DesiredPosition.SetVec(utils.Z_INDEX, r.DesiredPosition.At(utils.Z_INDEX, 0)+dist) // Increasing hanging the Z position

	return r.PositionError()
}

func (r *FlyingRobot) Move(dX, dY, dZ float64) float64 {
	r.DesiredPosition.SetVec(utils.X_INDEX, r.DesiredPosition.At(utils.X_INDEX, 0)+dX)
	r.DesiredPosition.SetVec(utils.Y_INDEX, r.DesiredPosition.At(utils.Y_INDEX, 0)+dY)
	r.DesiredPosition.SetVec(utils.Z_INDEX, r.DesiredPosition.At(utils.Z_INDEX, 0)+dZ)

	return r.PositionError()
}

func (r *FlyingRobot) Land() error {
	r.DesiredPosition.SetVec(utils.Z_INDEX, 0) // Setting the Z position to 0 to land the drone

	return nil
}
