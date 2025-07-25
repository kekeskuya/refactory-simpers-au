package lib

import (
	"dummy-simpers-au/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
	Size      int64 `json:"size"`
	Offset    int64 `json:"offset"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"total_page"`
	From      int64 `json:"from"`
	To        int64 `json:"to"`
}

// based on request by mas Haikal (24 Mei 2024).
// The response struct, json key name is defined such as below
type PaginationV2 struct {
	Total     int64 `json:"total"`
	TotalPage int64 `json:"lastPage"`
	Prev      int64 `json:"prevPage"`
	Next      int64 `json:"nextPage"`
	Size      int64 `json:"perPage"`
	Page      int64 `json:"currentPage"`
}

type HTTPError struct {
	Success bool   `json:"success" default:"false" `
	Message string `json:"message,omitempty"`
	Error   error  `json:"error"`
}

type APIResponse struct {
	Success bool        `json:"success" default:"true" `
	Message string      `json:"message,omitempty" example:"OK"`
	Data    interface{} `json:"data,omitempty"`
}

type APIResponsePaginated struct {
	APIResponse
	Pagination Pagination `json:"pagination,omitempty"`
}

// Return Response Error
// code int for http status
// message string for human readable
// err for error details
func RespondError(ctx *gin.Context, code int, message string, err error) {
	switch err {
	case constants.ErrorMessageUnauthorized:
		code = http.StatusUnauthorized
	case constants.ErrorMessageDataExists, constants.ErrorMessageUserAlreadyExists, constants.ErrorMessageTextTooLong, constants.ErrorMessageIdIsRequired, constants.ErrorMessageIdMustBeNumber, constants.ErrorMessageInvalidLimitValue, constants.ErrorMessageInvalidPageValue, constants.ErrorMessageUserAlreadyExists, constants.ErrorMessageInvalidInput:
		code = http.StatusBadRequest
	case constants.ErrorMessageDataNotFound, constants.ErrorMessageUserNotFound:
		code = http.StatusNotFound
	}

	ctx.JSON(code, HTTPError{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func RespondSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	if data == nil {
		data = constants.MessageSuccess
	}
	ctx.JSON(code, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func RespondSuccessPaginated(ctx *gin.Context, code int, message string, data interface{}, pagination Pagination) {
	ctx.JSON(code, APIResponsePaginated{
		APIResponse: APIResponse{
			Success: true,
			Message: message,
			Data:    data,
		},
		Pagination: pagination,
	})
}

const (
	MsgOk             = "OK"
	MsgDokumenSuccess = "Berhasil mendapatkan data %s"
)

// Error
const (
	ErrMissingRequired         = "missing data required"
	ErrBadPayload              = "please check your payload"
	ErrDeleteFailed            = "delete object failed"
	ErrForbidden               = "Forbidden"
	ErrUnauthorized            = "Unauthorized"
	ErrUserNotFound            = "user not found"
	ErrRecordNotFound          = "record not found"
	ErrWrongUsernameOrPassword = "wrong username or password"
)

const (
	ErrBucketDoesNotExist  = "storage: bucket doesn't exist"
	ErrStorageDoesNotExist = "storage: object doesn't exist"
)

const (
	ErrFileDoesNotExist = "file does not exist"
	ErrPathDoesNotExist = "path does not exist"
	ErrNoValidFiles     = "no valid files to be processed"
)

const (
	ErrPrepareQuery = "prepare query failed"
	ErrExecuteQuery = "execute query failed"
	ErrRunQuery     = "run query failed"
)
