package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

// Unit Tests for Tuna implementation of Prey interface
func TestTuna_GetSpeed(t *testing.T) {
	t.Run("speed is 0.0", func(t *testing.T) {
		// arrange
		impl := prey.NewTuna(0.0, nil)

		// act
		output := impl.GetSpeed()

		// assert
		outputSpeed := 0.0
		require.Equal(t, outputSpeed, output)
	})
	
	t.Run("speed is greater than 0.0", func(t *testing.T) {
		// arrange
		impl := prey.NewTuna(252.0, nil)

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
		impl := prey.NewTuna(0.0, nil)

		// act
		output := impl.GetPosition()

		// assert
		require.Nil(t, output)
	})
	
	t.Run("position is not nil", func(t *testing.T) {
		// arrange
		impl := prey.NewTuna(0.0, &positioner.Position{X: 0, Y: 0, Z: 0})

		// act
		output := impl.GetPosition()

		// assert
		require.NotNil(t, output)
	})
}