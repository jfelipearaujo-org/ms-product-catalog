package time_provider

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimeProvider(t *testing.T) {
	t.Run("Should return a new TimeProvider", func(t *testing.T) {
		// Arrange
		expected := &TimeProvider{}

		// Act
		timeProvider := NewTimeProvider(time.Now)

		// Assert
		assert.IsType(t, expected, timeProvider)
	})
}

func TestTimeProvider_GetTime(t *testing.T) {
	t.Run("Should return the current time", func(t *testing.T) {
		// Arrange
		timeProvider := NewTimeProvider(time.Now)

		// Act
		currentTime := timeProvider.GetTime()

		// Assert
		assert.NotEmpty(t, currentTime)
	})

	t.Run("Should return the correctly time", func(t *testing.T) {
		// Arrange
		customTime := "2024-04-21 20:40:11"

		expected, err := time.Parse("2006-01-02 15:04:05", customTime)
		assert.NoError(t, err)

		funcTime := func() time.Time {
			out, err := time.Parse("2006-01-02 15:04:05", customTime)
			if err != nil {
				panic(err)
			}
			return out
		}

		timeProvider := NewTimeProvider(funcTime)

		// Act
		currentTime := timeProvider.GetTime()

		// Assert
		assert.Equal(t, expected, currentTime)
	})
}
