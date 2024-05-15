package utils

import (
	"time"
)

func BuildSuccessfulResponse(success bool, timeStamp time.Time, data, metadata any) ResponseSuccess {
	return ResponseSuccess{
		Success:   success,
		Timestamp: timeStamp,
		Data:      data,
		MetaData:  metadata,
	}
}
