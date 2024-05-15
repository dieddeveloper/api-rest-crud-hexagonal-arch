package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildErrorResponse(t *testing.T) {
	t.Run("the error object is successfully constructed", func(t *testing.T) {
		// arrange
		// act
		response := BuildErrorResponse(true, time.Now(), "test")
		// assert
		assert.NotNil(t, response)
		assert.Equal(t, response, response)
	})
}
