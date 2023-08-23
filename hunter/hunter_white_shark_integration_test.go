package hunter

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

// Integration Test for White Shark implementation of Hunter
func TestIntegration_WhiteShark_Hunt(t *testing.T) {
	t.Run("Hunter can catch the prey - faster", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(p1, p2 *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		sm := simulator.NewCatchSimulatorDefault(100, ps)

		impl := &WhiteShark{
			speed:     10,
			position:  &positioner.Position{X: 0, Y: 0, Z: 0},
			simulator: sm,
		}

		// act
		inputPrey := prey.NewPreyStub()
		inputPrey.GetSpeedFunc = func() (speed float64) {speed = 5; return}
		inputPrey.GetPositionFunc = func() (position *positioner.Position) {
			position = &positioner.Position{X: 100, Y: 0, Z: 0}
			return
		}
		err := impl.Hunt(inputPrey)

		// assert
		outputErr := error(nil)
		require.ErrorIs(t, err, outputErr)
	})

	t.Run("Hunter can not catch the prey - faster but overtime", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(p1, p2 *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		sm := simulator.NewCatchSimulatorDefault(1, ps)

		impl := &WhiteShark{
			speed:     10,
			position:  &positioner.Position{X: 0, Y: 0, Z: 0},
			simulator: sm,
		}

		// act
		inputPrey := prey.NewPreyStub()
		inputPrey.GetSpeedFunc = func() (speed float64) {speed = 5; return}
		inputPrey.GetPositionFunc = func() (position *positioner.Position) {
			position = &positioner.Position{X: 100, Y: 0, Z: 0}
			return
		}
		err := impl.Hunt(inputPrey)

		// assert
		outputErr := ErrCanNotHunt
		outputErrMsg := "can not hunt the prey: shark can not catch the prey"
		require.ErrorIs(t, err, outputErr)
		require.EqualError(t, err, outputErrMsg)
	})

	t.Run("Hunter can not catch the prey - slower", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(p1, p2 *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		sm := simulator.NewCatchSimulatorDefault(100, ps)

		impl := &WhiteShark{
			speed:     5,
			position:  &positioner.Position{X: 0, Y: 0, Z: 0},
			simulator: sm,
		}

		// act
		inputPrey := prey.NewPreyStub()
		inputPrey.GetSpeedFunc = func() (speed float64) {speed = 10; return}
		inputPrey.GetPositionFunc = func() (position *positioner.Position) {
			position = &positioner.Position{X: 100, Y: 0, Z: 0}
			return
		}
		err := impl.Hunt(inputPrey)

		// assert
		outputErr := ErrCanNotHunt
		outputErrMsg := "can not hunt the prey: shark can not catch the prey"
		require.ErrorIs(t, err, outputErr)
		require.EqualError(t, err, outputErrMsg)
	})
}