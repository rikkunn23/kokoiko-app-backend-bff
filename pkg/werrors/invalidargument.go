package werrors

import "net/http"

type InvalidArgumentError struct {
	BaseError
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}
type InvalidArgumentErrorResponse struct {
	BaseErrorResponse
	// エラーごとに独自のフィールドを持たせたい場合はここに追加
}

// InvalidArgument ...
func InvalidArgument(err error) InvalidArgumentError {
	return InvalidArgumentError{
		BaseError: BaseError{
			Code:  http.StatusBadRequest,
			Err:   err,
			Stack: getStacks(),
		},
	}
}


// 以下はインターフェースを満たすために定義されている
// Error ...
func (e InvalidArgumentError) Error() string {
	return e.Err.Error()
}
// GetStacks ...
func (e InvalidArgumentError) GetStacks() []string {
	return e.Stack
}
