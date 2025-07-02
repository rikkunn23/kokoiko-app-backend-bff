INSERT INTO kokoiko_owner.tbf_user (f_user_no, f_email, f_tel, f_password, f_sei_name, f_mei_name, f_created_at, f_updated_at,f_sys_rec_dt, f_sys_rec_app, f_sys_rec_acnt,f_sys_upd_dt, f_sys_upd_app, f_sys_upd_acnt,f_del_flg)OVERRIDING SYSTEM VALUE
VALUES
    (1, 'yamada@example.com', '09012345678', '$2a$10$examplehash1','山田', '太郎',clock_timestamp(), clock_timestamp(),clock_timestamp(), 'seed_script', 'admin',clock_timestamp(), 'seed_script', 'admin', 0),
    (2, 'sato@example.com', '08098765432', '$2a$10$examplehash2','佐藤', '花子',clock_timestamp(), clock_timestamp(),clock_timestamp(), 'seed_script', 'admin',clock_timestamp(), 'seed_script', 'admin', 0);
