package utils

import "time"

type ResponseSuccess struct {
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
	Data      any       `json:"data"`
	MetaData  any       `json:"metadata,omitempty"`
}

type ResponseError struct {
	Success      bool      `json:"success"`
	Timestamp    time.Time `json:"timestamp"`
	ErrorMessage string    `json:"error"`
}
