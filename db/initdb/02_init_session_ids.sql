-- テーブル作成
CREATE TABLE session_ids (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL,
    session_id VARCHAR(255) UNIQUE NOT NULL
);