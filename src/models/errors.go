package models

import "errors"

// errors.go file contain a list of Sentinel Errors - specific errors that can be recognized across the different
// application layers.

var (
	// ErrorNotFound defines an error where a requested resource was not found, such as in a database call
	ErrorNotFound = errors.New("not found")
)
