package werrors

import (
	"net/http"
	"strings"
)

type ConflictError struct {
	BaseError
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

type ConflictErrorResponse struct {
	BaseErrorResponse
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

// Conflict ...
func Conflict(err error) ConflictError {
	return ConflictError{
		BaseError: BaseError{
			Code:  http.StatusConflict,
			Err:   err,
			Stack: getStacks(),
		},
	}
}

// 以下はインターフェースを満たすために定義されている
// Error ...
func (e ConflictError) Error() string {
	return strings.Replace(e.Err.Error(), "\n", ",", -1)
}
// GetStacks ...
func (e ConflictError) GetStacks() []string {
	return e.Stack
}
