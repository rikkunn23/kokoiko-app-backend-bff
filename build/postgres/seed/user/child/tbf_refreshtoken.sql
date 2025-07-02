INSERT INTO kokoiko_owner.tbf_refresh_token (
    f_user_no,
    f_refresh_token,
    f_expired_at,
    f_created_at,
    f_updated_at,
    f_sys_rec_dt,
    f_sys_rec_app,
    f_sys_rec_acnt,
    f_sys_upd_dt,
    f_sys_upd_app,
    f_sys_upd_acnt,
    f_del_flg
)
VALUES
    (
        1,
        'eyJfcmF0XzEiOiJzZWVkVG9rZW4xMTExMTExMTExIiwidHlwIjoiUmVmcmVzaFRva2VuIn0=', -- 例：base64のランダム文字列
        clock_timestamp() + interval '30 days',
        clock_timestamp(),
        clock_timestamp(),
        clock_timestamp(),
        'seed_script',
        'admin',
        clock_timestamp(),
        'seed_script',
        'admin',
        0
    ),
    (
        2,
        'eyJfcmF0XzIiOiJzZWVkVG9rZW4yMjIyMjIyMjIyIiwidHlwIjoiUmVmcmVzaFRva2VuIn0=',
        clock_timestamp() + interval '30 days',
        clock_timestamp(),
        clock_timestamp(),
        clock_timestamp(),
        'seed_script',
        'admin',
        clock_timestamp(),
        'seed_script',
        'admin',
        0
    );
