package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBuildErrorResponse(t *testing.T) {
	t.Run("the error object is successfully constructed", func(t *testing.T) {
		// arrange
		// act
		response := BuildErrorResponse(true, time.Now(), []string{"test"})
		// assert
		assert.NotNil(t, response)
		assert.Equal(t, response, response)
	})
}
