package werrors

import (
	"net/http"
	"strings"
)

type InternalError struct {
	BaseError
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

type InternalErrorResponse struct {
	BaseErrorResponse
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

// Internal ...
func Internal(err error) InternalError {
	return InternalError{
		BaseError: BaseError{
			Code:  http.StatusInternalServerError,
			Err:   err,
			Stack: getStacks(),
		},
	}
}

// 以下はインターフェースを満たすために定義されている

// Error ...
func (e InternalError) Error() string {
	return strings.Replace(e.Err.Error(), "\n", ",", -1)
}

// GetStacks ...
func (e InternalError) GetStacks() []string {
	return e.Stack
}
