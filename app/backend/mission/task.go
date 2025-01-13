package mission

import (
	"drone-rescue/robots"
)

type Task struct {
	// Name of the task
	Name string
	// Robots that are assigned to the task
	Robots []robots.Robot
	// Zones that the task involves
	Zones []Zone
}
