package hunter

import (
	"testdoubles/positioner"
	"testdoubles/prey"
)

// NewHunter return a mock implementation of Hunter
func NewHunterMock() *HunterMock {
	return &HunterMock{}
}

// Hunter is a mock implementation of Hunter
type HunterMock struct {
	HuntFunc func(pr prey.Prey) (err error)
	ConfigureFunc func(speed float64, position *positioner.Position)
	// observers
	Calls struct {
		Hunt int
		Configure int
	}
}

func (ht *HunterMock) Hunt(pr prey.Prey) (err error) {
	// observers
	ht.Calls.Hunt++

	err = ht.HuntFunc(pr)
	return
}

func (ht *HunterMock) Configure(speed float64, position *positioner.Position) {
	// observers
	ht.Calls.Configure++

	ht.ConfigureFunc(speed, position)
}