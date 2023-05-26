package core

import (
	"fmt"
)

var (
	ErrNotFound = func(msgs ...string) error {
		return fmt.Errorf("not found: %v", msgs)
	}

	ErrInvalid = func(msgs ...string) error {
		return fmt.Errorf("invalid: %v", msgs)
	}
)
