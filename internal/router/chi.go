package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rikkunn23/kokoiko-app-backend-bff/config"
	"github.com/rikkunn23/kokoiko-app-backend-bff/gen/api/master"
	mastercnt "github.com/rikkunn23/kokoiko-app-backend-bff/internal/controller/master"
	masterrepo "github.com/rikkunn23/kokoiko-app-backend-bff/internal/domain/master/repository"
	masterquery "github.com/rikkunn23/kokoiko-app-backend-bff/internal/domain/master/repository/query"
	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/postgres"
	masteruse "github.com/rikkunn23/kokoiko-app-backend-bff/internal/usecase/master"
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
	r.Get((config.Get().App.HealthCheckPath), healthCheckHandler)

	// mcli, err := postgres.NewMaster()
	// if err != nil {
	// 	log.Fatalf("failed to create master postgres client: %v", err)
	// }

	scli, err := postgres.NewSlave()
	if err != nil {
		log.Fatalf("failed to create slave postgres client: %v", err)
	}

	// マスター取得API
	master.HandlerFromMux(mastercnt.New(masteruse.New(masterrepo.New(scli, masterquery.New()))) ,r)
}
func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
  w.WriteHeader(http.StatusOK)
  _, _ = w.Write([]byte("OK"))
}
