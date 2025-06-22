package master

import (
	"net/http"

	apiMaster "github.com/rikkunn23/kokoiko-app-backend-bff/gen/api/master"
)

type IUsecase interface {
	CityList(r *http.Request, params apiMaster.CityListParams)(*apiMaster.ResponseCityList, error)
}
