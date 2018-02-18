package mta

import (
	"fmt"
)

const (
	FeedUndefinedErrorType ErrorType = iota
	DefaultErrorType
)

type ErrorType int

type FeedUndefinedError struct {
	e error
}

func (f FeedUndefinedError) Error() string {
	return f.e.Error()
}

type DefaultError struct {
	e error
}

func (f DefaultError) Error() string {
	return f.e.Error()
}

func newError(text string, et ErrorType) error {
	err := fmt.Errorf(text)
	switch et {
	case FeedUndefinedErrorType:
		return FeedUndefinedError{e: err}
	default:
		return DefaultError{e: err}
	}
}
