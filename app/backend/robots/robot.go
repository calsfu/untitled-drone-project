// Robot interface and BaseRobot struct
package robots

// import "fmt"

// Functions to be implemented by any robot
type Robot interface {
	// Description returns a description of the robot
	GetDescription() string
	// Run executes the robot's main functionality
	Run() error
	// Stop halts the robot's operation
	Stop() error
	// GetStatus returns the current status of the robot
	GetStatus() string
	// Name returns the name of the robot
	GetId() string
	// GetPosition returns the current position of the robot
	GetPosition() [3]float64
	// GetBatteryLevel returns the current battery level of the robot
	GetBatteryLevel() float64
}
