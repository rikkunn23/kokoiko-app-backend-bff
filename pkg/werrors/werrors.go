package werrors

import (
	"fmt"
	"runtime"
)

// IError ...
type IError interface {
	Error() string
	GetStacks() []string
}

// BaseError is an error struct
type BaseError struct {
	Code  int
	Err   error
	Stack []string
}

// BaseErrorResponse ...
type BaseErrorResponse struct {
	// Code ステータスコード
	// エラーコードを表記する
	Code int `json:"code"`

	// Msg ステータスメッセージ
	// エラー詳細は出力せずに固定文言を返す
	Msg string `json:"msg"`
}




//debug.PrintStack()を使うと、スタックトレースを取得できるが、
// そのままではファイル名と行番号が出力されないため、
// runtime.Callerを使ったカスタムスタックトレースを使ってスタックトレースを取得する
func getStacks() []string {
// 0	今の関数（runtime.Caller を呼んだ行）
// 1	1つ前に呼び出した関数
// 2	さらにその前に呼び出した関数
// だから２から12までの10
	i := 2
	stacks := []string{}
	// stackは決め打ちで10出しておく
	for i < 12 {
		// runtime.Callerはスタックトレースを取得する
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stacks = append(stacks, fmt.Sprintf("%s:%d", file, line))
		i++
	}
	return stacks
}
