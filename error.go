package types

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorString string

func (s ErrorString) Error() string {
	return string(s)
}

const (
	ErrNotExist         ErrorString = "not exist"
	ErrAlreadyExists    ErrorString = "already exists"
	ErrUnauthenticated  ErrorString = "unauthenticated"
	ErrPermissionDenied ErrorString = "permission denied"
)

type Error struct {
	code    int
	subCode int
	message string
}

type jsonError struct {
	Code    int    `json:"code,omitempty"`
	SubCode int    `json:"sub_code,omitempty"`
	Message string `json:"message,omitempty"`
}

var _ json.Marshaler = (*Error)(nil)
var _ json.Unmarshaler = (*Error)(nil)

func (e *Error) Code() int {
	return e.code
}

func (e *Error) SubCode() int {
	return e.subCode
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) Error() string {
	if e.message == "" {
		e.message = http.StatusText(e.code)
		if e.message == "" {
			e.message = fmt.Sprint(e.code)
		} else if e.subCode > 0 {
			e.message = fmt.Sprintf("%s (%d)", e.message, e.subCode)
		}
	}
	return e.message
}

func (e *Error) Is(target error) bool {
	if e == target {
		return true
	}

	if t, ok := target.(*Error); ok {
		return t.code == e.code && t.subCode == e.subCode && t.message == e.message
	}
	return false
}

func (e *Error) MarshalJSON() (text []byte, err error) {
	je := &jsonError{
		Code:    e.code,
		SubCode: e.subCode,
		Message: e.message,
	}
	return json.Marshal(je)
}

func (e *Error) UnmarshalJSON(text []byte) error {
	var je jsonError
	err := json.Unmarshal(text, &je)
	if err != nil {
		return err
	}
	e.code = je.Code
	e.subCode = je.SubCode
	e.message = je.Message
	return nil
}

func NewError(code int, format string, a ...any) *Error {
	if code <= 0 {
		panic("invalid code")
	}
	msg := fmt.Sprintf(format, a...)
	if msg == "" {
		msg = http.StatusText(code)
	}
	return &Error{
		code:    code,
		message: msg,
	}
}

func NewSubError(code, subCode int, message string) *Error {
	if code <= 0 {
		panic("invalid code")
	}

	if subCode <= 0 {
		panic("invalid subCode")
	}

	if message == "" {
		message = http.StatusText(code)
	}
	return &Error{
		code:    code,
		subCode: subCode,
		message: message,
	}
}