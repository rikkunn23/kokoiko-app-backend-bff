package config

import "time"

// AppConfig APIの設定情報
type AppConfig struct {
  Address         string `env:"GO_ADDRESS"`
  Port            string `env:"GO_PORT"`
  LogLevel        string `env:"GO_LOG_LEVEL"`
  LogEncoding     string `env:"GO_LOG_ENCODING"`
  RootDir         string `env:"GO_ROOT_DIR"`
  TimeZone        string `env:"TZ"`
  HealthCheckPath string `env:"GO_HEALTHCHECK_PATH"`
  Name            string `env:"GO_APP_NAME"`
  RecApp          string `env:"KOKOIKO_REC_APP"`
  RecAcnt         string `env:"KOKOIKO_REC_ACNT"`
  UpdApp          string `env:"KOKOIKO_UPD_APP"`
  UpdAcnt         string `env:"KOKOIKO_UPD_ACNT"`
  AgentDomain     string `env:"AGENT_DOMAIN"`
	JWTSecret		 		string `env:"JWT_SECRET"`
}

// PostgresConfig Postgresへの接続情報
type PostgresConfig struct {
  MasterHost   string `env:"POSTGRES_MASTER_HOST"`
  MasterPort   string `env:"POSTGRES_MASTER_PORT"`
  SlaveHost    string `env:"POSTGRES_SLAVE_HOST"`
  SlavePort    string `env:"POSTGRES_SLAVE_PORT"`
  Pass         string `env:"POSTGRES_PASSWORD"`
  User         string `env:"POSTGRES_USER"`
  DB           string `env:"POSTGRES_DB"`
  DebugMode    bool   `env:"POSTGRES_DEBUG_MODE"`
  MaxIdleConns int    `env:"POSTGRES_MAX_IDLE_CONNS"`
  MaxOpenConns int    `env:"POSTGRES_MAX_OPEN_CONNS"`
  SSLMode      string `env:"POSTGRES_SSL_MODE"`
  InLimit      int    `env:"POSTGRES_IN_LIMIT"`
  SelectLimit  int    `env:"POSTGRES_SELECT_LIMIT"`
}

// // AWSConfig ...
// type AWSConfig struct {
//   Region   string `env:"AWS_REGION"`
//   EndPoint string `env:"AWS_ENDPOINT"`
//   SQS      *SQSConfig
//   Cognito  *CognitoConfig
// }


// MS化するならここに設定する

// HTTPClientConfig ...
type HTTPClientConfig struct {
  TransportTimeout      time.Duration `env:"HTTP_CLIENT_TRANSPORT_TIMEOUT"`
  KeepAlive             time.Duration `env:"HTTP_CLIENT_KEEP_ALIVE"`
  MaxIdleConns          int           `env:"HTTP_CLIENT_MAX_IDLE_CONNS"`
  MaxIdleConnsPerHost   int           `env:"HTTP_CLIENT_MAX_IDLE_CONNS_PER_HOST"`
  MaxConnsPerHost       int           `env:"HTTP_CLIENT_MAX_CONNS_PER_HOST"`
  IdleConnsTimeout      time.Duration `env:"HTTP_CLIENT_IDLE_CONNS_TIMEOUT"`
  TLSHandShakeTimeout   time.Duration `env:"HTTP_CLIENT_TLS_HANDSHAKE_TIMEOUT"`
  ResponseHeaderTimeout time.Duration `env:"HTTP_CLIENT_RESPONSE_HEADER_TIMEOUT"`
  ExpectContinueTimeout time.Duration `env:"HTTP_CLIENT_EXPECT_CONTINUE_TIMEOUT"`
  Timeout               time.Duration `env:"HTTP_CLIENT_TIMEOUT"`
}


// Config 全ての設定情報を含む
type Config struct {
  App                        *AppConfig
  Postgres                   *PostgresConfig
  HTTPClient                 *HTTPClientConfig
}
