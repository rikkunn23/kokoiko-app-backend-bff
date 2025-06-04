package werrors

import (
	"net/http"
)

// UnauthorizedError ...
type UnauthorizedError struct {
	BaseError
	Detail *string
}

// UnauthorizedErrorResponse ...
type UnauthorizedErrorResponse struct {
	BaseErrorResponse
	// エラー毎に独自のフィールドを持たせたい場合はここに追加
	// 詳細メッセージ
	Detail *string `json:"detail,omitempty"`
}

// Unauthorized ..
func Unauthorized(err error) UnauthorizedError {
	return UnauthorizedError{
		BaseError: BaseError{
			Code:  http.StatusUnauthorized,
			Err:   err,
			Stack: getStacks(),
		},
	}
}

// Error ...
func (e UnauthorizedError) Error() string {
	return e.Err.Error()
}

// GetStacks ...
func (e UnauthorizedError) GetStacks() []string {
	return e.Stack
}
