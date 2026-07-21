-- +goose Up

-- Username must be unique
ALTER TABLE users
ADD CONSTRAINT users_username_key UNIQUE (username);

-- Username cannot be empty
ALTER TABLE users
ADD CONSTRAINT users_username_not_empty
CHECK (length(trim(username)) > 0);

-- Add new columns
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) NOT NULL UNIQUE DEFAULT (
    encode(sha256(random()::text::bytea),'hex')
),
ADD COLUMN name TEXT NOT NULL DEFAULT (
    'PRE_MIG_002_TEST_USER'
);

-- +goose Down

ALTER TABLE users
DROP COLUMN name,
DROP COLUMN api_key,
DROP CONSTRAINT users_username_not_empty,
DROP CONSTRAINT users_username_key;