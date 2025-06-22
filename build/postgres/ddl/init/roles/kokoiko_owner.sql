-- ロール作成
CREATE ROLE kokoiko_owner with PASSWORD 'kokoiko_owner' LOGIN;

-- サーチパス設定
ALTER ROLE kokoiko_owner SET search_path to kokoiko_owner;
