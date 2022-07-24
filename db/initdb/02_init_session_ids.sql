-- テーブル作成
CREATE TABLE session_ids (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    session_id INTEGER NOT NULL
);