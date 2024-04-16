package prey

import "testdoubles/positioner"

// Constructor
// NewPreyStub creates a new Prey
func NewPreyStub() (prey *PreyStub) {
	prey = &PreyStub{}
	return
}

// Stub
// PreyStub is a stub for Prey
type PreyStub struct {
	// Add functions of the Prey interface
	GetPositionFunc func() (position *positioner.Position)
	GetSpeedFunc    func() (speed float64)
}

// GetPosition calls the GetPositionFunc
func (p *PreyStub) GetPosition() (position *positioner.Position) {
	return p.GetPositionFunc()
}

// GetSpeed calls the GetSpeedFunc
func (p *PreyStub) GetSpeed() (speed float64) {
	return p.GetSpeedFunc()
}
