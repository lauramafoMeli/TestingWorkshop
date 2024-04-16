package positioner

// Constructor
// NewPositionerStub creates a new Positioner
func NewPositionerStub() (positioner *PositionerStub) {
	positioner = &PositionerStub{}
	return
}

// Stub
// PositionerStub is a stub for Positioner
type PositionerStub struct {
	// Add functions of the Positioner interface
	GetLinearDistanceFunc func(from, to *Position) (linearDistance float64)
}

// GetLinearDistance calls the GetLinearDistanceFunc
func (p *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	return p.GetLinearDistanceFunc(from, to)
}
