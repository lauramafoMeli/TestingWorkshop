package simulator

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

// Unit Tests for CatchSimulatorDefault
func TestCatchSimulatorDefault_CanCatch(t *testing.T) {
	t.Run("Hunter can catch the prey - hunter faster", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		cfgImpl := &ConfigCatchSimulatorDefault{MaxTimeToCatch: 100, Positioner: ps}
		impl := NewCatchSimulatorDefault(cfgImpl)

		// act
		inputHunter := &Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &Subject{Speed: 5, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
		output := impl.CanCatch(inputHunter, inputPrey)

		// assert
		outputOk := true
		require.Equal(t, outputOk, output)
	})
	
	t.Run("Hunter can not catch the prey - hunter faster but long distance", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 1000
			return
		}

		cfgImpl := &ConfigCatchSimulatorDefault{MaxTimeToCatch: 100, Positioner: ps}
		impl := NewCatchSimulatorDefault(cfgImpl)

		// act
		inputHunter := &Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &Subject{Speed: 5, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
		output := impl.CanCatch(inputHunter, inputPrey)

		// assert
		outputOk := false
		require.Equal(t, outputOk, output)
	})

	t.Run("Hunter can not catch the prey - hunter slower", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		cfgImpl := &ConfigCatchSimulatorDefault{MaxTimeToCatch: 100, Positioner: ps}
		impl := NewCatchSimulatorDefault(cfgImpl)

		// act
		inputHunter := &Subject{Speed: 5, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &Subject{Speed: 10, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
		output := impl.CanCatch(inputHunter, inputPrey)

		// assert
		outputOk := false
		require.Equal(t, outputOk, output)
	})
}