package positioner_test

import (
	. "testdoubles/positioner"
	"testing"
)

// Test of the GetLinearDistance function
func TestGetLinearDistance(t *testing.T) {
	// Positions are negative
	t.Run("Positions are negative", func(t *testing.T) {
		// Given
		// Create a new PositionerDefault
		positioner := NewPositionerDefault()
		position1 := Position{X: -1, Y: -1, Z: -1}
		position2 := Position{X: -2, Y: -2, Z: -2}

		// When
		// Call the GetLinearDistance function
		linearDistance := positioner.GetLinearDistance(&position1, &position2)

		// Then
		// Check the result
		if linearDistance != 1.7320508075688772 {
			t.Errorf("The linear distance is not correct")
		}
	})
	// Positions are positive
	t.Run("Positions are positive", func(t *testing.T) {
		// Given
		// Create a new PositionerDefault
		positioner := NewPositionerDefault()
		position1 := Position{X: 1, Y: 1, Z: 1}
		position2 := Position{X: 2, Y: 2, Z: 2}

		// When
		// Call the GetLinearDistance function
		linearDistance := positioner.GetLinearDistance(&position1, &position2)

		// Then
		// Check the result
		if linearDistance != 1.7320508075688772 {
			t.Errorf("The linear distance is not correct")
		}
	})
	// Positions linear distance do not return a decimal number
	t.Run("Positions linear distance do not return a decimal number", func(t *testing.T) {
		// Given
		// Create a new PositionerDefault
		positioner := NewPositionerDefault()
		position1 := Position{X: 1, Y: 1, Z: 1}
		position2 := Position{X: 1, Y: 1, Z: 2}

		// When
		// Call the GetLinearDistance function
		linearDistance := positioner.GetLinearDistance(&position1, &position2)

		// Then
		// Check the result
		if linearDistance != 1 {
			t.Errorf("The linear distance is not correct")
		}
	})
}
