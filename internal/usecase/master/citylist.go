package master

import (
	"net/http"

	apiMaster "github.com/rikkunn23/kokoiko-app-backend-bff/gen/api/master"
)


func (u *Usecase) CityList(r *http.Request, param apiMaster.CityListParams) (*apiMaster.ResponseCityList, error) {
	ctx := r.Context()

	cityRes, err := u.repo.SelectCity(ctx, param.TdfkCd)
	if err != nil {
		return nil, err
	}
	if cityRes == nil {
		return &apiMaster.ResponseCityList{}, nil
	}

	res := make([]apiMaster.ResponseCity, len(cityRes))
	for i, city := range cityRes {
		res[i] = apiMaster.ResponseCity{
			CityCd:   city.CityCd,
			CityName: city.CityName.String,
			CityKana: city.CityKana.String,
			Bms:      city.Bms.Int64,
			Lms:      city.Lms.Int64,
			TdfkCd:   city.TdfkCd,
			TdfkName: city.TdfkName.String,

		}
	}
	return &apiMaster.ResponseCityList{
		Results	: res,
	}, nil
}
