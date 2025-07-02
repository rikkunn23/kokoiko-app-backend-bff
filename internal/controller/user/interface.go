package user

import (
	"net/http"

	"github.com/rikkunn23/kokoiko-app-backend-bff/gen/api/user"
)

type IUsecase interface {
	CreateUser(r *http.Request, reqBody user.RequestUserCreate) (user.ResponseUserCreate, error)
}
