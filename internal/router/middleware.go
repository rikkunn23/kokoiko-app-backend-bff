package router

import (
	"errors"
	"net/http"

	"github.com/rikkunn23/kokoiko-app-backend-bff/config"
	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/controller"
	"github.com/rikkunn23/kokoiko-app-backend-bff/pkg/werrors"
)

func verifyAPIKey(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// ヘルスチェックのパスはAPIキーの検証をスキップする
    if r.URL.Path == config.Get().App.HealthCheckPath {
			next.ServeHTTP(w, r)
      return
    }
		// APIキーの検証をここで行う(デフォルトはAPI_KEY=api-key)
    if r.Header.Get("X-Api-Key") != config.GetAPIKey() {
			// Error()を実装していないとエラー型として扱えない
			var err error = werrors.Unauthorized(errors.New("invalid api key"))
      // err := werrors.Unauthorized(errors.New("invalid api key"))
      controller.ResponseHandler(r.Context(), w, nil, err)
      return
    }
		// 複数のミドルウェアやハンドラーをつなげて処理するため他のハンドラーに処理を渡す
    next.ServeHTTP(w, r.WithContext(r.Context()))
  }
	// この関数をhttp.Handlerインターフェースを満たす型に変換してHTTPミドルウェアとして返す
  return http.HandlerFunc(fn)
}
