package entity

type CreateUser struct {
	Email     string `db:"f_email"`
	Tel       string `db:"f_tel"`
	Password  string `db:"f_password"`
	SeiName   string `db:"f_sei_name"`
	MeiName   string `db:"f_mei_name"`
	RecApp    string `db:"f_sys_rec_app"`
	RecAcnt   string `db:"f_sys_rec_acnt"`
	UpdApp    string `db:"f_sys_upd_app"`
	UpdAcnt   string `db:"f_sys_upd_acnt"`
}
