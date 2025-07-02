package user

import (
	"context"
	"time"

	userent "github.com/rikkunn23/kokoiko-app-backend-bff/internal/domain/user/entity"
)

type IRepository interface {
	CreateUser(ctx context.Context, entity userent.CreateUser ) (int64, error)
	SaveRefreshToken(ctx context.Context, userNo int64, refreshTokem string, expiresAt time.Time) error
}
