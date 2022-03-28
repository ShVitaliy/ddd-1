package ddd

import (
	"errors"
	"fmt"
	"testing"
)

func TestDomainErrorCreation(t *testing.T) {
	errorConstructors := map[string]func(string, ...any) error{
		"domain":            func(format string, a ...any) error { return NewDomainError(format, a...) },
		"invalid state":     func(format string, a ...any) error { return NewInvalidStateError(format, a...) },
		"invalid argument":  func(format string, a ...any) error { return NewInvalidArgumentError(format, a...) },
		"entity not found":  func(format string, a ...any) error { return NewEntityNotFoundError(format, a...) },
		"duplicate entity":  func(format string, a ...any) error { return NewDuplicateEntityError(format, a...) },
		"not authenticated": func(format string, a ...any) error { return NewNotAuthenticatedError(format, a...) },
		"access denied":     func(format string, a ...any) error { return NewAccessDeniedError(format, a...) },
	}

	for errorName, errorConstructor := range errorConstructors {
		err := errorConstructor("this is %s error", errorName)
		if err.Error() != fmt.Sprintf("this is %s error", errorName) {
			t.Errorf("Failed to create %s error.", errorName)
		}
	}
}

func TestErrorToDomainErrorConversion(t *testing.T) {
	errorConverters := map[string]func(error) error{
		"domain":            func(err error) error { return ErrorToDomainError(err) },
		"invalid state":     func(err error) error { return ErrorToInvalidStateError(err) },
		"invalid argument":  func(err error) error { return ErrorToInvalidArgumentError(err) },
		"entity not found":  func(err error) error { return ErrorToEntityNotFoundError(err) },
		"duplicate entity":  func(err error) error { return ErrorToDuplicateEntityError(err) },
		"not authenticated": func(err error) error { return ErrorToNotAuthenticatedError(err) },
		"access denied":     func(err error) error { return ErrorToAccessDeniedError(err) },
	}

	for errorName, errorConverter := range errorConverters {
		err2err := errorConverter(errors.New(errorName))
		if err2err.Error() != errorName {
			t.Errorf("Failed to convert error to %s error.", errorName)
		}
	}
}
