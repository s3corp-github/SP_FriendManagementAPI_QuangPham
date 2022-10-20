BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(100) NOT NULL UNIQUE,
    phone      VARCHAR(20),
    is_active  BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS relations
(
    id         SERIAL PRIMARY KEY,
    requester_id     INTEGER NOT NULL REFERENCES users(id),
    addressee_id     INTEGER NOT NULL REFERENCES users(id),
    requester_email     VARCHAR(100) NOT NULL,
    addressee_email     VARCHAR(100) NOT NULL,
    relation_type INTEGER DEFAULT 0,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMIT;



