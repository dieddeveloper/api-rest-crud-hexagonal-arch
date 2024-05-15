package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBuildSuccessfulResponse(t *testing.T) {
	t.Run("response object is successfully built", func(t *testing.T) {
		// arrange
		// act
		response := BuildSuccessfulResponse(true, time.Now(), "test", "test")
		// assert
		assert.NotNil(t, response)
		assert.Equal(t, response, response)
	})
}
