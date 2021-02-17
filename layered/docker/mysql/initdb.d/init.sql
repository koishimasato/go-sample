CREATE DATABASE IF NOT EXISTS sample_db;

CREATE TABLE IF NOT EXISTS users
(
    id        VARCHAR(255) NOT NULL,
    name      VARCHAR(32)  NOT NULL,
    user_type INT          NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (id, name, user_type) VALUES ('1', 'taro', 1);
INSERT INTO users (id, name, user_type) VALUES ('2', 'jiro', 2);
INSERT INTO users (id, name, user_type) VALUES ('3', 'saburo', 1);
