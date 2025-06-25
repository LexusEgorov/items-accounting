package models

import "fmt"

var (
	ErrUnique     = fmt.Errorf("already exists")
	ErrNotFound   = fmt.Errorf("not found")
	ErrNotUpdated = fmt.Errorf("not updated")
)
