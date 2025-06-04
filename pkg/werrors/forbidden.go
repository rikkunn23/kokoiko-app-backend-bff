package werrors

import "net/http"

type ForbiddenError struct {
	BaseError
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

type ForbiddenErrorResponse struct {
	BaseErrorResponse
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

// Forbidden ...
func Forbidden(err error) ForbiddenError {
	return ForbiddenError{
		BaseError: BaseError{
			Code:  http.StatusForbidden,
			Err:   err,
			Stack: getStacks(),
		},
	}
}

// 以下はインターフェースを満たすために定義されている
// Error ...
func (e ForbiddenError) Error() string {
	return e.Err.Error()
}
// GetStacks ...
func (e ForbiddenError) GetStacks() []string {
	return e.Stack
}
