package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	userent "github.com/rikkunn23/kokoiko-app-backend-bff/internal/domain/user/entity"
	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/postgres"
	"github.com/rikkunn23/kokoiko-app-backend-bff/pkg/werrors"
)

type Repository struct {
	masterClient *postgres.Client
	query	IQuery
}

func New(client *postgres.Client, query IQuery) *Repository {
	repo := &Repository{
		masterClient: client,
		query:  query,
	}
	return repo
}

// CreateUser...
func (r *Repository) CreateUser(ctx context.Context, user userent.CreateUser) (int64, error) {
	namedStmt, err := r.masterClient.DB.PrepareNamedContext(ctx, r.query.CreateUser())
	if err != nil {
		return 0, werrors.Internal(errors.Errorf("ユーザ登録 %s", err))
	}
	defer namedStmt.Close()
	var userNo int64
	err = namedStmt.GetContext(ctx, &userNo, user)
  if err != nil {
    return 0, werrors.Internal(errors.Errorf("ユーザ登録 GetContext 失敗: %s", err))
  }
	fmt.Println("userNo:", userNo)
	return userNo, nil
}

// SaveRefreshToken...
func (r *Repository) SaveRefreshToken(ctx context.Context, userNo int64, refreshToken string, expiresAt time.Time) error {
	namedStmt, err := r.masterClient.DB.PrepareNamedContext(ctx, r.query.SaveRefreshToken())
	if err != nil {
		return werrors.Internal(errors.Errorf("リフレッシュトークン保存 %s", err))
	}
	defer namedStmt.Close()

	arg := map[string]interface{}{
		"user_no":       userNo,
		"refresh_token": refreshToken,
		"expired_at":    expiresAt,
	}

	if _, err := namedStmt.ExecContext(ctx, arg); err != nil {
		return werrors.Internal(errors.Errorf("リフレッシュトークン保存 ExecContext 失敗: %s", err))
	}
	return nil
}
