package query

type Query struct {
	// 共通的に持たせたいパラメーターがある場合は記載
}

func New() *Query {
	return new(Query)
}

// SelectCity ...
func (q *Query) SelectCity() string {
	return `
SELECT
	f_del_flg AS f_del_flg
	, f_tdfk_cd AS f_tdfk_cd
	, f_city_cd AS f_city_cd
	, f_tdfk_name AS f_tdfk_name
	, f_city_name AS f_city_name
	, f_bms AS f_bms
	, f_lms AS f_lms
	, f_city_kana AS f_city_kana
FROM
	kokoiko_owner.mbf_city
WHERE
	f_tdfk_cd = :tdfk_cd
	AND
	f_del_flg = 0
ORDER BY
	f_city_cd
`
}
