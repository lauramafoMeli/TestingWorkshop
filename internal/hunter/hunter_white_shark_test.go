package hunter_test

import (
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation of the Hunter interface
func TestHunterWhiteShark_Hunt(t *testing.T) {
	t.Run("white shark hunts a prey - has speed and short distance", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return true
		}

		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     100,
			Position:  &positioner.Position{X: 1, Y: 1, Z: 1},
			Simulator: sm,
		})

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := error(nil)
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunt a prey - has short speed", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return false
		}

		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     1,
			Position:  &positioner.Position{X: 1, Y: 1, Z: 1},
			Simulator: sm,
		})

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := hunter.ErrCanNotHunt; expErrMsg := "can not hunt the prey: shark can not catch the prey"
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.EqualError(t, err, expErrMsg)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunt a prey - has long distance", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return false
		}

		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     100,
			Position:  &positioner.Position{X: 1000, Y: 1000, Z: 1000},
			Simulator: sm,
		})

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := hunter.ErrCanNotHunt; expErrMsg := "can not hunt the prey: shark can not catch the prey"
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.EqualError(t, err, expErrMsg)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})
}

func TestHunterWhiteShark_Configure(t *testing.T) {
	t.Run("set speed to 100", func(t *testing.T) {
		// arrange
		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     0,
			Position:  nil,
			Simulator: nil,
		})

		// act
		inputSpeed := 100.0
		inputPosition := (*positioner.Position)(nil)
		impl.Configure(inputSpeed, inputPosition)

		// assert
		// outputSpeed := 100.0
		// outputPosition := (*positioner.Position)(nil)
		// require.Equal(t, outputSpeed, impl.speed)
		// require.Equal(t, outputPosition, impl.position)
	})

	t.Run("set position to (1, 2, 3)", func(t *testing.T) {
		// arrange
		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     0,
			Position:  nil,
			Simulator: nil,
		})

		// act
		inputSpeed := 0.0
		inputPosition := &positioner.Position{X: 1, Y: 2, Z: 3}
		impl.Configure(inputSpeed, inputPosition)

		// assert
		// outputSpeed := 0.0
		// outputPosition := &positioner.Position{X: 1, Y: 2, Z: 3}
		// require.Equal(t, outputSpeed, impl.speed)
		// require.Equal(t, outputPosition, impl.position)
	})
}