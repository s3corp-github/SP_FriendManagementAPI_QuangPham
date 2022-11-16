BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    email      TEXT NOT NULL UNIQUE,
    name       TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS user_friends
(
    id                  SERIAL PRIMARY KEY,
    requester_id        INTEGER NOT NULL REFERENCES users(id),
    target_id           INTEGER NOT NULL REFERENCES users(id),
    relation_type       INTEGER DEFAULT 1 , -- 1:FRIEND, 2:SUBSCRIPTION, 3:BLOCKED
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
    );

COMMIT;
