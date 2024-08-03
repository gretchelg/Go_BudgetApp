package random

import (
	"github.com/lithammer/shortuuid"
)

// GenerateRandomUUID returns a random string of the given length
func GenerateRandomUUID(length int) string {
	result := shortuuid.New()

	// guard against index overflow
	if length >= len(result) {
		return result
	}

	// truncate the result to the specified length
	truncatedResult := result[0:length]

	return truncatedResult
}
