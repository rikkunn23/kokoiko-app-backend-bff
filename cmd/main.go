package main

import (
	"log"
	"net"
	"net/http"

	"github.com/rikkunn23/kokoiko-app-backend-bff/config"
	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/router"
)
func main() {

	// 環境変数の読み込み
  err := config.NewConfig()
  if err != nil {
    log.Fatalf("failed to init config: %v", err)
  }
  // wslog.New(true, config.SlogLevel())
  // if err != nil {
  //   log.Fatalf("failed to create aws config: %v", err)
  // }

  // if err := newrelic.Initialize(); err != nil {
  //   log.Fatalf("failed to init newrelic: %v\n", err)
  // }

  l, err := net.Listen("tcp", config.Get().App.Address+":"+config.Get().App.Port)
  if err != nil {
    log.Fatal(err)
  }
  if err := http.Serve(l, router.CreateServiceRouter()); err != nil {
    log.Fatal(err)
  }
  log.Println("start")
}
