package hunter

import (
	"errors"
	"testdoubles/positioner"
	"testdoubles/prey"
)

// Hunter is an interface that represents a hunter
type Hunter interface {
	// Hunt hunts the prey
	Hunt(prey prey.Prey) (err error)
	// Configure configures the hunter
	Configure(speed float64, position *positioner.Position)
}

var (
	ErrCanNotHunt = errors.New("can not hunt the prey")
)