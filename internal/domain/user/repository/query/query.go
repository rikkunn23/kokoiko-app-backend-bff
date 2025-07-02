package query

type Query struct {
	// 共通的に持たせたいパラメーターがある場合は記載
}

func New() *Query {
	return new(Query)
}

// CreateUser ...
func (q *Query) CreateUser() string {
	return `INSERT INTO kokoiko_owner.tbf_user (
	f_email,
	f_tel,
	f_password,
	f_sei_name,
	f_mei_name,
	f_sys_rec_app,
	f_sys_rec_acnt,
	f_sys_upd_app,
	f_sys_upd_acnt
) VALUES (
	:f_email,
	:f_tel,
	:f_password,
	:f_sei_name,
	:f_mei_name,
	:f_sys_rec_app,
	:f_sys_rec_acnt,
	:f_sys_upd_app,
	:f_sys_upd_acnt
)
	RETURNING f_user_no
	`
}

// SaveRefreshToken ...
func (q *Query) SaveRefreshToken() string {
return `INSERT INTO kokoiko_owner.tbf_refresh_token (
	f_user_no,
	f_refresh_token,
	f_expired_at
) VALUES (
	:user_no,
	:refresh_token,
	:expired_at
)`
}
