package errors

import "fmt"

const (
	ErrStorageNil    = "storage not initialized"
	ErrTaskExists    = "task with ID %d already exist"
	ErrTaskNotFound  = "task with ID %d not found"
	ErrIndexNotFound = "index %d not found"
	ErrInsertingTask = "error inserting task"
)

type MyErrors struct {
	Code   int
	Msg    string
	Method string
}

func (m *MyErrors) Error() string {
	return fmt.Sprintf("method: %s, code: %d, message: %s", m.Method, m.Code, m.Msg)
}
