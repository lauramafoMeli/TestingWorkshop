package simulator

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

// Integration Test for CatchSimulatorDefault with the Default Implementation of Positioner
func TestIntegration_CatchSimulatorDefault_CanCatch(t *testing.T) {
	t.Run("Hunter can catch the prey - hunter faster", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerDefault()

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
		ps := positioner.NewPositionerDefault()

		cfgImpl := &ConfigCatchSimulatorDefault{MaxTimeToCatch: 100, Positioner: ps}
		impl := NewCatchSimulatorDefault(cfgImpl)

		// act
		inputHunter := &Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &Subject{Speed: 5, Position: &positioner.Position{X: 1000, Y: 0, Z: 0}}
		output := impl.CanCatch(inputHunter, inputPrey)

		// assert
		outputOk := false
		require.Equal(t, outputOk, output)
	})

	t.Run("Hunter can not catch the prey - hunter slower", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerDefault()

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