DROP TABLE IF EXISTS post CASCADE;
CREATE TABLE IF NOT EXISTS post(
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES author(id),
    content TEXT NOT NULL,
    delete_flag CHAR(1) DEFAULT '0'
);

DROP TABLE IF EXISTS author CASCADE;
CREATE TABLE IF NOT EXISTS author(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    delete_flag CHAR(1) DEFAULT '0'
);

DROP TABLE IF EXISTS comment CASCADE;
CREATE TABLE IF NOT EXISTS comment(
    id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES post(id),
    author_id INTEGER REFERENCES author(id),
    content TEXT NOT NULL,
    delete_flag CHAR(1) DEFAULT '0'
);