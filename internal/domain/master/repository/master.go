package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/domain/master/entity"
	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/postgres"
	"github.com/rikkunn23/kokoiko-app-backend-bff/pkg/werrors"
)

type Repository struct {
	client *postgres.Client
	query	IQuery
}

func New(client *postgres.Client, query IQuery) *Repository {
	repo := &Repository{
		client: client,
		query:  query,
	}
	return repo
}




// SelectCity ...
func (r *Repository) SelectCity(ctx context.Context, tdfkCd string) ([]entity.SelectCityResult, error) {
	namedStmt, err := r.client.DB.PrepareNamedContext(ctx, r.query.SelectCity())
	if err != nil {
		return nil, werrors.Internal(errors.Errorf("tbf_city取得 %s", err))
	}
	defer namedStmt.Close()

	arg := map[string]interface{}{
		"tdfk_cd": tdfkCd,
	}
	var res []entity.SelectCityResult
	if err := namedStmt.SelectContext(ctx, &res, arg); err != nil {
		return nil, werrors.Internal(errors.Errorf("tbf_city取得 %s", err))
	}
	return res, nil
}
