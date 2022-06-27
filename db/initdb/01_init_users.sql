-- テーブル作成
CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- サンプルレコード作成
INSERT INTO users (name, email, password) VALUES('test1', 'test1@test.com', 'test1');
INSERT INTO users (name, email, password) VALUES('test2', 'test2@test.com', 'test2');
INSERT INTO users (name, email, password) VALUES('test3', 'test3@test.com', 'test3');