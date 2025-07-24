package constants

import "errors"

var (
	ErrorMessageUnauthorized      error = errors.New("unauthorized")
	ErrorMessageUserNotFound      error = errors.New("user not found")
	ErrorMessageUserAlreadyExists error = errors.New("user exists")
	ErrorMessageDataNotFound      error = errors.New("data not found")
	ErrorMessageDataExists        error = errors.New("data exists")
	ErrorMessageTextTooLong       error = errors.New("text is too long")
	ErrorMessageIdIsRequired      error = errors.New("id is required")
	ErrorMessageIdMustBeNumber    error = errors.New("id must be a number")
	ErrorMessageInvalidLimitValue error = errors.New("invalid limit value")
	ErrorMessageInvalidPageValue  error = errors.New("invalid page value")
	ErrorMessageInvalidInput      error = errors.New("invalid input")
)
