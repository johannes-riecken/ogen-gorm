// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type Error string

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/user
type User struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *User) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *User) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *User) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *User) SetName(val string) {
	s.Name = val
}

// UsersPostCreated is response for UsersPost operation.
type UsersPostCreated struct{}
