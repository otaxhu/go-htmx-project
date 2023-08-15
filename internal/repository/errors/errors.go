package errors

import "fmt"

var (
	ErrNoRows = fmt.Errorf("no rows were found in the database")
)
