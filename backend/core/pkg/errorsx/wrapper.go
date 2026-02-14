package errorsx

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
)

func New(message string) error {
	_, file, line, _ := runtime.Caller(1)

	return &Error{
		stack: []StackFrame{{
			File:    filepath.Base(file),
			Line:    line,
			Message: message,
		}},
		httpStatus: http.StatusInternalServerError,
	}
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	_, file, line, _ := runtime.Caller(1)

	var e *Error
	if errors.As(err, &e) {
		e.stack = append(e.stack, StackFrame{
			File:    filepath.Base(file),
			Line:    line,
			Message: message,
		})
		return e
	}

	return &Error{
		stack: []StackFrame{
			{
				File:    filepath.Base(file),
				Line:    line,
				Message: err.Error(),
			},
			{
				File:    filepath.Base(file),
				Line:    line,
				Message: message,
			},
		},
		httpStatus: http.StatusInternalServerError,
	}
}

func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

func Extract(err error) error {
	if err == nil {
		return nil
	}

	var e *Error
	if errors.As(err, &e) {
		return e
	}

	_, file, line, _ := runtime.Caller(1)
	return &Error{
		stack: []StackFrame{{
			File:    filepath.Base(file),
			Line:    line,
			Message: err.Error(),
		}},
		httpStatus: http.StatusInternalServerError,
	}
}
