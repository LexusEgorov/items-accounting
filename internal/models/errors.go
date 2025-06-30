package models

import (
	"fmt"
)

var (
	ErrUnique                = fmt.Errorf("already exists")
	ErrNotFound              = fmt.Errorf("not found")
	ErrNotUpdated            = fmt.Errorf("not updated")
	ErrConfigPathNotProvided = fmt.Errorf("config path didn't provide")
	ErrBadConfigPort         = fmt.Errorf("port must be upper than 0")
	ErrBadResponseTime       = fmt.Errorf("response time must be upper than 0ms")
	ErrBadUserName           = fmt.Errorf("username is required")
	ErrBadPassword           = fmt.Errorf("password is required")
	ErrBadDBName             = fmt.Errorf("db name is required")
	ErrMigrationsNotProvided = fmt.Errorf("migrations path didn't provide")
)

func NewEmptyErr(field string) error {
	return fmt.Errorf("field '%s' is required!", field)
}
