package werrors

import "net/http"

type NotFoundError struct {
	BaseError
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

type NotFoundErrorResponse struct {
	BaseErrorResponse
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

// NotFound ...
func NotFound(err error) NotFoundError {
	return NotFoundError{
		BaseError: BaseError{
			Code:  http.StatusNotFound,
			Err:   err,
			Stack: getStacks(),
		},
	}
}

// 以下はインターフェースを満たすために定義されている
// Error ...
func (e NotFoundError) Error() string {
	return e.Err.Error()
}
// GetStacks ...
func (e NotFoundError) GetStacks() []string {
	return e.Stack
}
