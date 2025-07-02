package user

import (
	"encoding/json"
	"net/http"

	"github.com/rikkunn23/kokoiko-app-backend-bff/gen/api/user"
	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/controller"
)

// Controller interface for master
type Controller struct {
	usecase IUsecase
}

// NewController ...
func New(usecase IUsecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

// UserPost ...
func (c *Controller) UserPost(w http.ResponseWriter, r *http.Request, params user.UserPostParams) {
	var reqBody user.RequestUserCreate
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		controller.ResponseHandler(r.Context(), w, nil, err)
		return
	}
	res, err := c.usecase.CreateUser(r, reqBody)
		controller.ResponseHandler(r.Context(), w, res, err)

		// TODOステータスコードを201で返す
}
