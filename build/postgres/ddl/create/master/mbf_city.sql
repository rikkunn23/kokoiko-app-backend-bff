-- ------------ Write CREATE-TABLE-stage scripts -----------
CREATE TABLE IF NOT EXISTS kokoiko_owner.mbf_city(
    f_sys_rec_dt timestamp(0) with time zone not null default (clock_timestamp())::timestamp(0) with time zone,
    f_sys_rec_app varchar(100) not null default 'Unknown',
    f_sys_rec_acnt varchar(100) not null default 'Unknown',
    f_sys_upd_dt timestamp(0) with time zone not null default (clock_timestamp())::timestamp(0) with time zone,
    f_sys_upd_app varchar(100) not null default 'Unknown',
    f_sys_upd_acnt varchar(100) not null default 'Unknown',
    f_del_flg smallint not null default (0)::smallint,
    f_tdfk_cd varchar(2) not null,
    f_city_cd varchar(3) not null,
    f_tdfk_name varchar(30),
    f_city_name varchar(50),
    f_bms bigint,
    f_lms bigint,
    f_city_kana varchar(100)
)
        WITH (
    fillfactor = 90
        );
-- ------------ Write CREATE-CONSTRAINT-stage scripts -----------
ALTER TABLE kokoiko_owner.mbf_city DROP CONSTRAINT IF EXISTS mbf_city_pkey;

ALTER TABLE kokoiko_owner.mbf_city
ADD CONSTRAINT mbf_city_pkey PRIMARY KEY (f_tdfk_cd, f_city_cd);

ALTER INDEX kokoiko_owner.mbf_city_pkey SET (deduplicate_items = off);
-- -----------------------comments--------------------------
COMMENT ON TABLE kokoiko_owner.mbf_city IS '市区町村';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_sys_rec_dt IS 'レコード登録日時';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_sys_rec_app IS 'レコードを登録したアプリ';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_sys_rec_acnt IS 'レコードを登録したアカウント';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_sys_upd_dt IS 'レコード更新日時';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_sys_upd_app IS 'レコードを更新したアプリ';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_sys_upd_acnt IS 'レコードを更新したアカウント';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_del_flg IS '論理削除フラグ';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_tdfk_cd IS '都道府県コード';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_city_cd IS '市区町村コード';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_tdfk_name IS '都道府県名';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_city_name IS '市区町村名';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_bms IS '位置情報緯度(ミリ秒)';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_lms IS '位置情報経度(ミリ秒)';
COMMENT ON COLUMN kokoiko_owner.mbf_city.f_city_kana IS '読み仮名（ひらがな）';
