package simulator

// Constructor
// NewCatchSimulatorMock creates a new Simulator
func NewCatchimulatorMock() (simulator *CatchSimulatorMock) {
	simulator = &CatchSimulatorMock{}
	return
}

// Mock
// CatchSimulatorMock is a mock for CatchSimulator
type CatchSimulatorMock struct {
	// Add functions of the CatchSimulator interface
	CanCatchFunc func(hunter, prey *Subject) (canCatch bool)

	// Observer struct
	Calls struct {
		// Register the n times the method is called
		CanCatch int
	}
}

// CanCatch calls the CanCatchFunc
func (s *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	// Register the call
	s.Calls.CanCatch++

	return s.CanCatchFunc(hunter, prey)
}
