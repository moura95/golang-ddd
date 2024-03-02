package util

import "errors"

var (
	ErrorBadRequest     = errors.New("Bad Request: One or more fields are invalid")
	ErrorBadRequestUuid = errors.New("Bad Request!! Uuid: Invalid")
	ErrorDatabaseCreate = errors.New("Error creating record")
	ErrorDatabaseRead   = errors.New("Error reading record")
	ErrorDatabaseUpdate = errors.New("Error updating record")
	ErrorDatabaseDelete = errors.New("Error deleting record")
)
