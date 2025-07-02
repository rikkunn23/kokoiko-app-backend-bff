-- ------------ Write CREATE-TABLE-stage scripts -----------
CREATE TABLE IF NOT EXISTS kokoiko_owner.tbf_user (
    f_user_no        bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    f_email          varchar(255) not null unique, -- メールアドレス（unique制約）
    f_tel   varchar(20) not null unique, -- 電話番号（unique制約）
    f_password  varchar(255) not null,
    f_sei_name       varchar(100) not null, -- 姓
    f_mei_name       varchar(100) not null, -- 名
    f_created_at     timestamp(0) with time zone not null default (clock_timestamp())::timestamp(0) with time zone,
    f_updated_at     timestamp(0) with time zone not null default (clock_timestamp())::timestamp(0) with time zone,
    f_sys_rec_dt   timestamp(0) with time zone not null default (clock_timestamp())::timestamp(0) with time zone,
    f_sys_rec_app  varchar(100) not null default 'Unknown',
    f_sys_rec_acnt varchar(100) not null default 'Unknown',
    f_sys_upd_dt   timestamp(0) with time zone not null default (clock_timestamp())::timestamp(0) with time zone,
    f_sys_upd_app  varchar(100) not null default 'Unknown',
    f_sys_upd_acnt varchar(100) not null default 'Unknown',
    f_del_flg      smallint not null default (0)::smallint
)
WITH (
    fillfactor = 90
);

COMMENT ON TABLE kokoiko_owner.tbf_user IS 'ユーザー情報';
COMMENT ON COLUMN kokoiko_owner.tbf_user.f_user_no IS 'ユーザーNo（主キー。他テーブル参照用）';
COMMENT ON COLUMN kokoiko_owner.tbf_user.f_email IS 'メールアドレス';
COMMENT ON COLUMN kokoiko_owner.tbf_user.f_tel IS '電話番号';
COMMENT ON COLUMN kokoiko_owner.tbf_user.f_password IS 'パスワード(ハッシュ)';
COMMENT ON COLUMN kokoiko_owner.tbf_user.f_sei_name IS '姓';
COMMENT ON COLUMN kokoiko_owner.tbf_user.f_mei_name IS '名';
COMMENT ON COLUMN kokoiko_owner.tbf_user.f_del_flg IS '論理削除フラグ';
