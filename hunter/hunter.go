package hunter

import (
	"errors"
	"testdoubles/prey"
)

// Hunter is an interface that represents a hunter
type Hunter interface {
	// Hunt hunts the prey
	Hunt(prey prey.Prey) (err error)
}

var (
	ErrCanNotHunt = errors.New("can not hunt the prey")
)