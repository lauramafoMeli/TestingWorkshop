package prey

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

// Unit Tests for Tuna implementation of Prey interface
func TestTuna_GetSpeed(t *testing.T) {
	t.Run("speed is 0", func(t *testing.T) {
		// arrange
		impl := &Tuna{speed: 0.0, position: nil}

		// act
		output := impl.GetSpeed()

		// assert
		outputSpeed := 0.0
		require.Equal(t, outputSpeed, output)
	})
	
	t.Run("speed is greater than 0", func(t *testing.T) {
		// arrange
		impl := &Tuna{speed: 252.0, position: nil}

		// act
		output := impl.GetSpeed()

		// assert
		outputSpeed := 252.0
		require.Equal(t, outputSpeed, output)
	})
}

func TestTuna_GetPosition(t *testing.T) {
	t.Run("position is nil", func(t *testing.T) {
		// arrange
		impl := &Tuna{speed: 0, position: nil}

		// act
		output := impl.GetPosition()

		// assert
		require.Nil(t, output)
	})
	
	t.Run("position is not nil", func(t *testing.T) {
		// arrange
		impl := &Tuna{speed: 0, position: &positioner.Position{X: 0, Y: 0, Z: 0}}

		// act
		output := impl.GetPosition()

		// assert
		require.NotNil(t, output)
	})
}

func TestTuna_Configure(t *testing.T) {
	t.Run("set speed to 100", func(t *testing.T) {
		// arrange
		impl := &Tuna{speed: 0.0, position: nil}

		// act
		inputSpeed := 100.0
		inputPosition := (*positioner.Position)(nil)
		impl.Configure(inputSpeed, inputPosition)

		// assert
		outputSpeed := 100.0
		outputPosition := (*positioner.Position)(nil)
		require.Equal(t, outputSpeed, impl.speed)
		require.Equal(t, outputPosition, impl.position)
	})

	t.Run("set position to (100, 0, 0)", func(t *testing.T) {
		// arrange
		impl := &Tuna{speed: 0, position: nil}

		// act
		inputSpeed := 0.0
		inputPosition := &positioner.Position{X: 100, Y: 0, Z: 0}
		impl.Configure(inputSpeed, inputPosition)

		// assert
		outputSpeed := 0.0
		outputPosition := &positioner.Position{X: 100, Y: 0, Z: 0}
		require.Equal(t, outputSpeed, impl.speed)
		require.Equal(t, outputPosition, impl.position)
	})
}