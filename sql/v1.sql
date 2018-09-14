-- init db
DROP DATABASE IF EXISTS depsea;
CREATE DATABASE depsea;
\c depsea;

-- init tables
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL,
    nickname VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    headimg TEXT DEFAULT '',
    sex INT DEFAULT 1,
    mobile VARCHAR(11) DEFAULT 1,
    status INT DEFAULT 0,
    create_time TIMESTAMP DEFAULT now(),
    update_time TIMESTAMP DEFAULT now()
);

DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    description VARCHAR(256) DEFAULT '',
    status INT DEFAULT 0,
    create_time TIMESTAMP DEFAULT now(),
    update_time TIMESTAMP DEFAULT now()
);

DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    description VARCHAR(256) DEFAULT '',
    status INT DEFAULT 0,
    create_time TIMESTAMP DEFAULT now(),
    update_time TIMESTAMP DEFAULT now()
);

DROP TABLE IF EXISTS articles;
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    description VARCHAR(256) DEFAULT '',
    category INT NOT NULL,
    tags INT[] NOT NULL,
    body TEXT DEFAULT '',
    create_user INT NOT NULL,
    status INT DEFAULT 0,
    create_time TIMESTAMP DEFAULT now(),
    update_time TIMESTAMP DEFAULT now(),
    publish_time TIMESTAMP DEFAULT now()
);

-- init default data
INSERT INTO users (username, nickname, password) VALUES ('admin', 'admin', 'e10adc3949ba59abbe56e057f20f883e');