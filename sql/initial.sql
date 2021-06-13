DROP TABLE IF EXISTS forums CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS threads CASCADE;

CREATE UNLOGGED TABLE IF NOT EXISTS users
(
    id       SERIAL PRIMARY KEY,
    nickname CITEXT UNIQUE NOT NULL,
    fullname TEXT          NOT NULL,
    about    TEXT,
    email    CITEXT UNIQUE
);

CREATE UNLOGGED TABLE IF NOT EXISTS forums
(
    id           SERIAL PRIMARY KEY,
    title        TEXT                                                 NOT NULL,
    userNickname CITEXT REFERENCES users (nickname) ON DELETE CASCADE NOT NULL,
    slug         CITEXT UNIQUE,
    posts        INTEGER,
    threads      INTEGER
);

CREATE UNLOGGED TABLE IF NOT EXISTS threads
(
    id      SERIAL PRIMARY KEY,
    title   TEXT,
    author  CITEXT REFERENCES users (nickname) ON DELETE CASCADE NOT NULL,
    forum   CITEXT REFERENCES forums (slug) ON DELETE CASCADE,
    message TEXT, -- описание ветки
    votes   INTEGER DEFAULT 0                                    NOT NULL,
    slug    TEXT                                                 NOT NULL,
    created TIMESTAMP with time zone
);

CREATE UNLOGGED TABLE IF NOT EXISTS posts
(
    id       SERIAL PRIMARY KEY,
    parent   INTEGER REFERENCES posts (id) DEFAULT NULL,
    forum    CITEXT REFERENCES forums (slug) ON DELETE CASCADE    NOT NULL,
    author   CITEXT REFERENCES users (nickname) ON DELETE CASCADE NOT NULL,
    thread   INTEGER REFERENCES threads (id) ON DELETE CASCADE    NOT NULL,
    created  TIMESTAMP with time zone,
    message  TEXT,
    isEdited BOOLEAN                       DEFAULT FALSE
);

select *
from users;
select *
from forums;
select *
from threads;

select *
from posts;