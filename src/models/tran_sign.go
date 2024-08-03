package models

import (
	"fmt"
)

const (
	// TranSignCredit defines the credit tran sign
	TranSignCredit TranSign = "CR"

	// TranSignDebit defines the dedit tran sign
	TranSignDebit TranSign = "DR"
)

// TranSign defines the type for Transaction Sign, in order to limit allowable values
type TranSign string

// Validate validates the value and returns error if invalid.
func (t TranSign) Validate() error {
	switch t {
	case TranSignCredit:
		return nil
	case TranSignDebit:
		return nil
	default:
		return fmt.Errorf("unknown TranSign, expected CR or DR only, but got %s", t)
	}
}
