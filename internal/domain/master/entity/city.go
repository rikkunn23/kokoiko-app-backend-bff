package entity

import "database/sql"


type SelectCityResult struct {
	DelFlg string `db:"f_del_flg"`
	TdfkCd string `db:"f_tdfk_cd"`
	CityCd string `db:"f_city_cd"`
	TdfkName sql.NullString `db:"f_tdfk_name"`
	CityName sql.NullString `db:"f_city_name"`
	Bms sql.NullInt64 `db:"f_bms"`
	Lms sql.NullInt64 `db:"f_lms"`
	CityKana sql.NullString `db:"f_city_kana"`
}
