package logger

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rikkunn23/kokoiko-app-backend-bff/pkg/werrors"
	"github.com/rikkunn23/kokoiko-app-backend-bff/pkg/wslog"
)

type ctxKey int
// コンテキスト用キー
const (
  MethodKey ctxKey = iota
  RequestURIKey
  TaskIDKey
  StackKey
  TraceKey
  SpanKey
  UserNoKey
  UserPhoneNumberKey
  UserAccountIDKey
  CognitoUserIDKey
  StandbyUserKey
  StandbyJobKey
  StandbyYukichiKey
  // log type
  LogType = "bff"
)

// ErrorWithCtx ...
func ErrorWithCtx(ctx context.Context, args ...interface{}) {
  args, stack := getArgsWithStacks(args...)
  ctx = context.WithValue(ctx, StackKey, stack)
  // errorレベルのログは通知を飛ばせるように監視用のフィールドを追加
  p := getWithParam(ctx)
  p = append(p, "notification", true)
  wslog.Get().ErrorContext(ctx, getMessage("", args), p...)
}

// ErrorfWithCtx ...
func ErrorfWithCtx(ctx context.Context, format string, args ...interface{}) {
  args, stack := getArgsWithStacks(args...)
  ctx = context.WithValue(ctx, StackKey, stack)
  p := getWithParam(ctx)
  p = append(p, "notification", true)
  wslog.Get().ErrorContext(ctx, getMessage(format, args), p...)
}

// WinfoWithCtx ...
func WinfoWithCtx(ctx context.Context, err werrors.IError) {
  stack := getWerrStacks(err)
  ctx = context.WithValue(ctx, StackKey, stack)
  wslog.Get().InfoContext(ctx, getWerrMessage("", err), getWithParam(ctx)...)
}

// WwarnWithCtx ...
func WwarnWithCtx(ctx context.Context, err werrors.IError) {
  stack := getWerrStacks(err)
  ctx = context.WithValue(ctx, StackKey, stack)
  wslog.Get().WarnContext(ctx, getWerrMessage("", err), getWithParam(ctx)...)
}


// WerrorWithCtx ...
func WerrorWithCtx(ctx context.Context, err werrors.IError) {
  stack := getWerrStacks(err)
  ctx = context.WithValue(ctx, StackKey, stack)
  p := getWithParam(ctx)
  p = append(p, "notification", true)
  wslog.Get().ErrorContext(ctx, getWerrMessage("", err), p...)
}


// ログメッセージをテンプレートと引数のエラーを組み合わせて作成する
func getMessage(template string, fmtArgs []interface{}) string {
  if len(fmtArgs) == 0 {
    return template
  }
  if template != "" {
    return fmt.Sprintf(template, fmtArgs...)
  }
  if len(fmtArgs) == 1 {
    if str, ok := fmtArgs[0].(string); ok {
      return str
    }
  }
  return fmt.Sprint(fmtArgs...)
}

// 1個だけなら「エラー分解」してメッセージとスタックを分ける
// 2個以上なら「そのまま」扱う（複雑な分解せずgetMessageなど他の関数でまとめて文字列化する）
func getArgsWithStacks(args ...interface{}) ([]interface{}, string) {
  if len(args) != 1 {
    return args, ""
  }
  var stack string
  switch args[0].(type) {
  case error:
    ss := strings.Split(fmt.Sprintf("%+v", args[0]), "\n")
    args[0] = ss[0]
    stack = strings.Join(ss[1:], "\n")
  default:
  }
  return args, stack
}

// IErrorからメッセージを取得するしてテンプレートと組み合わせてログメッセージを作成する
func getWerrMessage(template string, err werrors.IError) string {
  if template != "" {
    return fmt.Sprintf(template, err.Error())
  }
  return fmt.Sprint(err.Error())
}

// IErrorからスタックトレースを取得してログに出力する
func getWerrStacks(err werrors.IError) string {
  return strings.Join(err.GetStacks(), "\n")
}

// getWithParam ... プロジェクトで共通の追加のログ項目を返す
func getWithParam(ctx context.Context) []interface{} {
  params := []interface{}{
    "server_time", time.Now(),
    "log_type", LogType,
    "request_id", ctx.Value(middleware.RequestIDKey),
    "method", ctx.Value(MethodKey),
    "request_uri", ctx.Value(RequestURIKey),
    "error.stack", ctx.Value(StackKey),
    "user_no", ctx.Value(UserNoKey),
    "user_account_id", ctx.Value(UserAccountIDKey),

		// 認証はcognitoを使わないのでコメントアウト
    // "cognito_user_id", ctx.Value(CognitoUserIDKey),

		// こちらにMSを追加する場合追加する
    // "standby_user", ctx.Value(StandbyUserKey),

  }

// newrelicのトランザクションがコンテキストにある場合は、トランザクションのメタデータを追加する

  // if txn := newrelic.FromContext(ctx); txn != nil {
  //   md := txn.GetLinkingMetadata()
  //   params = append(params,
  //     logcontext.KeyTraceID, md.TraceID,
  //     logcontext.KeySpanID, md.SpanID,
  //     logcontext.KeyEntityName, md.EntityName,
  //     logcontext.KeyEntityType, md.EntityType,
  //     logcontext.KeyEntityGUID, md.EntityGUID,
  //     logcontext.KeyHostname, md.Hostname,
  //   )
  // }
  return params
}
