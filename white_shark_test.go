package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// arrange
		// - white shark
		ws := hunt.NewWhiteShark(true, false, 100)
		
		// act
		tuna := hunt.NewTuna("A", 50.0)
		err := ws.Hunt(tuna)

		// assert
		require.NoError(t, err)
		require.False(t, ws.Hungry)
		require.True(t, ws.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// arrange
		// - white shark
		ws := hunt.NewWhiteShark(false, false, 100)

		// act
		tuna := hunt.NewTuna("A", 50.0)
		err := ws.Hunt(tuna)

		// assert
		require.ErrorIs(t, err, hunt.ErrSharkIsNotHungry)
		require.EqualError(t, err, hunt.ErrSharkIsNotHungry.Error())
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// arrange
		// - white shark
		ws := hunt.NewWhiteShark(true, true, 100)

		// act
		tuna := hunt.NewTuna("A", 50.0)
		err := ws.Hunt(tuna)

		// assert
		require.ErrorIs(t, err, hunt.ErrSharkIsTired)
		require.EqualError(t, err, hunt.ErrSharkIsTired.Error())
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// arrange
		// - white shark
		ws := hunt.NewWhiteShark(true, false, 50)

		// act
		tuna := hunt.NewTuna("A", 100.0)
		err := ws.Hunt(tuna)

		// assert
		require.ErrorIs(t, err, hunt.ErrSharkIsSlower)
		require.EqualError(t, err, hunt.ErrSharkIsSlower.Error())
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// arrange
		// - white shark
		ws := hunt.NewWhiteShark(true, false, 100)

		// act
		err := ws.Hunt(nil)

		// assert
		require.ErrorIs(t, err, hunt.ErrTunaIsNil)
		require.EqualError(t, err, hunt.ErrTunaIsNil.Error())
	})
}