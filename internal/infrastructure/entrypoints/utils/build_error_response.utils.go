package utils

import (
	"time"
)

func BuildErrorResponse(success bool, timeStamp time.Time, err string) ResponseError {
	return ResponseError{
		Success:      success,
		Timestamp:    timeStamp,
		ErrorMessage: err,
	}
}
