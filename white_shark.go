package hunt

import (
	"errors"
)

var (
	// ErrSharkIsTired indicates that the shark is tired
	ErrSharkIsTired = errors.New("can not hunt, shark is tired")
	// ErrSharkIsNotHungry indicates that the shark is not hungry
	ErrSharkIsNotHungry = errors.New("can not hunt, shark is not hungry")
	// ErrSharkIsSlower indicates that the shark could not catch the prey
	ErrSharkIsSlower = errors.New("can not hunt, shark is slower than the prey")
)

// NewWhiteShark creates a new WhiteShark
func NewWhiteShark(hungry bool, tired bool, speed float64) (w *WhiteShark) {
	w = &WhiteShark{
		hungry: hungry,
		tired:  tired,
		speed:  speed,
	}
	return
}

// WhiteShark is an implementation of the Hunter interface
type WhiteShark struct {
	// hungry indicates if the shark is hungry
	hungry bool
	// tired indicates if the shark is tired
	tired bool
	// speed indicates the speed of the shark
	speed float64
}

func (w *WhiteShark) Hunt(tuna *Tuna) (err error) {
	// check if the shark can hunt
	if !w.hungry {
		err = ErrSharkIsNotHungry
		return
	}
	if w.tired {
		err = ErrSharkIsTired
		return
	}
	if w.speed < tuna.speed {
		err = ErrSharkIsSlower
		return
	}

	// hunt done
	w.hungry = false
	w.tired = true
	return
}