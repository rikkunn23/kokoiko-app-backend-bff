CREATE TABLE IF NOT EXISTS kokoiko_owner.tbf_refresh_token (
    f_user_no          bigint NOT NULL,
    f_refresh_token    varchar(512) NOT NULL, -- リフレッシュトークン
    f_expired_at       timestamp(0) with time zone NOT NULL, -- 有効期限
    f_created_at       timestamp(0) with time zone NOT NULL DEFAULT (clock_timestamp())::timestamp(0) with time zone,
    f_updated_at       timestamp(0) with time zone NOT NULL DEFAULT (clock_timestamp())::timestamp(0) with time zone,
    f_sys_rec_dt       timestamp(0) with time zone NOT NULL DEFAULT (clock_timestamp())::timestamp(0) with time zone,
    f_sys_rec_app      varchar(100) NOT NULL DEFAULT 'Unknown',
    f_sys_rec_acnt     varchar(100) NOT NULL DEFAULT 'Unknown',
    f_sys_upd_dt       timestamp(0) with time zone NOT NULL DEFAULT (clock_timestamp())::timestamp(0) with time zone,
    f_sys_upd_app      varchar(100) NOT NULL DEFAULT 'Unknown',
    f_sys_upd_acnt     varchar(100) NOT NULL DEFAULT 'Unknown',
    f_del_flg          smallint NOT NULL DEFAULT (0)::smallint,
    PRIMARY KEY (f_user_no),
    FOREIGN KEY (f_user_no) REFERENCES kokoiko_owner.tbf_user(f_user_no)
)
WITH (
    fillfactor = 90
);

COMMENT ON TABLE kokoiko_owner.tbf_refresh_token IS 'リフレッシュトークン管理テーブル';
COMMENT ON COLUMN kokoiko_owner.tbf_refresh_token.f_user_no IS 'ユーザーNo（主キー。他テーブル参照）';
COMMENT ON COLUMN kokoiko_owner.tbf_refresh_token.f_refresh_token IS 'リフレッシュトークン（署名なしのランダム文字列）';
COMMENT ON COLUMN kokoiko_owner.tbf_refresh_token.f_expired_at IS 'リフレッシュトークンの有効期限';
