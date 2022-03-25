package ddd

import "fmt"

// --------------------------------------------------------------------------------------------------------------------
// DOMAIN ERROR
// --------------------------------------------------------------------------------------------------------------------

type DomainError struct {
	error
	err string
}

func (e DomainError) Error() string {
	return e.err
}

func ErrorToDomainError(err error) DomainError {
	return DomainError{err: err.Error()}
}

func NewDomainError(format string, a ...interface{}) DomainError {
	return DomainError{err: fmt.Sprintf(format, a...)}
}

// --------------------------------------------------------------------------------------------------------------------
// INVALID STATE ERROR
// --------------------------------------------------------------------------------------------------------------------

type InvalidStateError struct {
	DomainError
}

func ErrorToInvalidStateError(err error) InvalidStateError {
	return InvalidStateError{ErrorToDomainError(err)}
}

func NewInvalidStateError(format string, a ...interface{}) InvalidStateError {
	return InvalidStateError{NewDomainError(format, a...)}
}

// --------------------------------------------------------------------------------------------------------------------
// INVALID ARGUMENT ERROR
// --------------------------------------------------------------------------------------------------------------------

type InvalidArgumentError struct {
	DomainError
}

func ErrorToInvalidArgumentError(err error) InvalidArgumentError {
	return InvalidArgumentError{ErrorToDomainError(err)}
}

func NewInvalidArgumentError(format string, a ...interface{}) InvalidArgumentError {
	return InvalidArgumentError{NewDomainError(format, a...)}
}

// --------------------------------------------------------------------------------------------------------------------
// ENTITY NOT FOUND ERROR
// --------------------------------------------------------------------------------------------------------------------

type EntityNotFoundError struct {
	DomainError
}

func ErrorToEntityNotFoundError(err error) EntityNotFoundError {
	return EntityNotFoundError{ErrorToDomainError(err)}
}

func NewEntityNotFoundError(format string, a ...interface{}) EntityNotFoundError {
	return EntityNotFoundError{NewDomainError(format, a...)}
}

// --------------------------------------------------------------------------------------------------------------------
// DUPLICATE ENTITY ERROR
// --------------------------------------------------------------------------------------------------------------------

type DuplicateEntityError struct {
	DomainError
}

func ErrorToDuplicateEntityError(err error) DuplicateEntityError {
	return DuplicateEntityError{ErrorToDomainError(err)}
}

func NewDuplicateEntityError(format string, a ...interface{}) DuplicateEntityError {
	return DuplicateEntityError{NewDomainError(format, a...)}
}

// --------------------------------------------------------------------------------------------------------------------
// NOT AUTHENTICATED ERROR
// --------------------------------------------------------------------------------------------------------------------

type NotAuthenticatedError struct {
	DomainError
}

func ErrorToNotAuthenticatedError(err error) NotAuthenticatedError {
	return NotAuthenticatedError{ErrorToDomainError(err)}
}

func NewNotAuthenticatedError(format string, a ...interface{}) NotAuthenticatedError {
	return NotAuthenticatedError{NewDomainError(format, a...)}
}

// --------------------------------------------------------------------------------------------------------------------
// ACCESS DENIED ERROR
// --------------------------------------------------------------------------------------------------------------------

type AccessDeniedError struct {
	DomainError
}

func ErrorToAccessDeniedError(err error) AccessDeniedError {
	return AccessDeniedError{ErrorToDomainError(err)}
}

func NewAccessDeniedError(format string, a ...interface{}) AccessDeniedError {
	return AccessDeniedError{NewDomainError(format, a...)}
}
