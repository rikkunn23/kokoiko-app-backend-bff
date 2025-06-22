package master

import (
	"context"

	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/domain/master/entity"
)


type IRepository interface {
	SelectCity(ctx context.Context, TdfkCd string) ([]entity.SelectCityResult, error)
}
