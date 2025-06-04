package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// CreateServiceRouter uses for production
func CreateServiceRouter() *chi.Mux {
  r := chi.NewRouter()
  r.Use(render.SetContentType(render.ContentTypeJSON))
  // ミドルウェアの追加はここ
  r.Use(middleware.RequestID)
  r.Use(verifyAPIKey)

	// 待機系への接続は一旦コメントアウト
  // r.Use(addContext)
  // r.Use(writeLogFinish)
  // r.Use(writeRequestLog)
  // r.Use(handlePanic)

  applyRoutes(r)
  return r
}
// 新しいserviceができたらここに追加する
func applyRoutes(r *chi.Mux) {
	r.Get("/", healthCheckHandler)
}
func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
  w.WriteHeader(http.StatusOK)
  _, _ = w.Write([]byte("OK"))
}
