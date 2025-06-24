package models

import "fmt"

const (
	ErrUniqueCode = "23505"
)

var (
	ErrUnique     = fmt.Errorf("already exists")
	ErrNotFound   = fmt.Errorf("not found")
	ErrNotUpdated = fmt.Errorf("not updated")
)
