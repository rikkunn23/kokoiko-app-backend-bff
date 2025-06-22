package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	// ポスグレを使うためにはドライバのインポートが必要
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"

	"github.com/rikkunn23/kokoiko-app-backend-bff/config"
	"github.com/rikkunn23/kokoiko-app-backend-bff/pkg/werrors"
)

// postgres driver
const (
  driver string = "postgres"
)

// Client ...
type Client struct {
  DB *sqlx.DB
}
// マスタ接続はデータの書き込み(insert, update, delete)
var mcli *Client
// スレーブ接続はデータの読み取り(select)
var scli *Client


// NewMaster ..
func NewMaster() (*Client, error) {
  if mcli != nil {
    return mcli, nil
  }
  var src string
  posConf := config.Get().Postgres

  // 環境変数をコンテナに入れ込むところで空の値については挙動が不安定になるので、
  // posConf.Passが空の場合はpassword=%sの部分を明記しない。
  src = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s TimeZone=%s",
    posConf.MasterHost, posConf.MasterPort, posConf.User, posConf.DB, posConf.SSLMode, config.Get().App.TimeZone)

  var err error
  mcli, err = newClient(src)
  if err != nil {
    return nil, err
  }
  return mcli, nil
}


// NewSlave ..
func NewSlave() (*Client, error) {
  if scli != nil {
    return scli, nil
  }
  var src string
  posConf := config.Get().Postgres
  // 環境変数をコンテナに入れ込むところで空の値については挙動が不安定になるので、
  // posConf.Passが空の場合はpassword=%sの部分を明記しない。
  src = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s TimeZone=%s",
    posConf.SlaveHost, posConf.SlavePort, posConf.User, posConf.DB, posConf.SSLMode, config.Get().App.TimeZone)
  var err error
  scli, err = newClient(src)
  if err != nil {
    return nil, err
  }
  return scli, nil
}



func newClient(src string) (*Client, error) {
  posConf := config.Get().Postgres

  if posConf.Pass != "" {
    src += fmt.Sprintf(" password=%s", posConf.Pass)
  }

  db, err := sqlx.Open(driver, src)
  if err != nil {
    return nil, werrors.Internal(errors.Errorf("failed to connect database: %s", err))
  }

	// 接続プール内のアイドルとして保持できる接続数を設定
  db.SetMaxIdleConns(posConf.MaxIdleConns)
	// 接続プール内の同時接続の最大接続数を設定
  db.SetMaxOpenConns(posConf.MaxOpenConns)
  client := &Client{
    DB: db,
  }
  return client, nil
}
// Close ..
func (c *Client) Close() {
  c.DB.Close()
}

// // NewTest ..
// func NewTest(name, src string) (*Client, error) {
//   if config.IsAWSEnv() {
//     return nil, werrors.Internal(errors.New("failed to create test client: not test env"))
//   }
//   posConf := config.Get().Postgres
//   if posConf.Pass != "" {
//     src += fmt.Sprintf(" password=%s", posConf.Pass)
//   }
//   db, err := sqlx.Open(name, src)
//   if err != nil {
//     return nil, werrors.Internal(errors.Errorf("failed to connect database: %s", err))
//   }
//   db.SetMaxIdleConns(posConf.MaxIdleConns)
//   db.SetMaxOpenConns(posConf.MaxOpenConns)
//   client := &Client{
//     DB: db,
//   }
//   return client, nil
// }
