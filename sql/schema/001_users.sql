-- +goose Up
-- below here are up statemtns

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
-- below here are down statments

DROP TABLE users;