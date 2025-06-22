package master

import (
	"net/http"

	apiMaster "github.com/rikkunn23/kokoiko-app-backend-bff/gen/api/master"
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

// CityList ...
func (c *Controller) CityList(w http.ResponseWriter, r *http.Request, params apiMaster.CityListParams) {
	res, err := c.usecase.CityList(r, params)
	controller.ResponseHandler(r.Context(), w, res, err)
}
