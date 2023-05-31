BEGIN;

CREATE TABLE users
(
    id         uuid PRIMARY KEY,
    username   VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT NOW()
);

CREATE TYPE credential_type AS ENUM ('password');

CREATE TABLE credentials
(
    user_id    uuid            NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    type       credential_type NOT NULL,
    password   VARCHAR(255)    NOT NULL,
    created_at TIMESTAMP       NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, type)
);

CREATE TABLE tokens
(
    hash            BYTEA PRIMARY KEY,
    user_id         UUID         NOT NULL,
    scope           VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP    NOT NULL,
    last_visited_at TIMESTAMP    NOT NULL
);

COMMIT;
