package cni

import (
	"encoding/json"
	"fmt"
)

// Error represents the object we return to runtime in the case of a failure
type Error struct {
	Version string `json:"cniVersion"`
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Details string `json:"details,omitempty"`
}

//NewError generates a new CNIError
func NewError(code int, message string) *Error {
	return &Error{
		Version: CNIVersion,
		Code:    code,
		Message: message,
	}
}

//NewDetailedError generates a new CNIError with details
func NewDetailedError(code int, message, details string) *Error {
	e := NewError(code, message)
	e.Details = details
	return e
}

//Marshal marshals the error into a json byte array
func (e *Error) Marshal() []byte {
	ebytes, err := json.Marshal(e)
	if err != nil {
		return []byte(fmt.Sprintf("{\"cniVersion\": \"%v\", \"code\": 99, \"msg\":\"error marshaling error\", \"details\":\"there was an error marshaling the original error\"}", CNIVersion))
	}

	return ebytes
}

//PrepareExit returns the error code and json bytes to be printed to stdout by the calling program
func PrepareExit(err error, code int, message string) (int, []byte) {
	if err == nil {
		return code, NewError(code, message).Marshal()
	}
	return code, NewDetailedError(code, message, err.Error()).Marshal()
}
