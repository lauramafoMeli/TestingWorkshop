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