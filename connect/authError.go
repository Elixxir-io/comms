package connect

import (
	"errors"
	"gitlab.com/elixxir/primitives/id"
	"strings"
)

const baseAuthErr = "Failed to authenticate"

// AuthError returns a valid authorization error on the given id
func AuthError(id *id.ID) error {
	return errors.New(baseAuthErr + " id: " + id.String())
}

// IsAuthError returns true if the passed error is a valid auth error
func IsAuthError(err error) bool {
	return strings.Contains(err.Error(), baseAuthErr)
}