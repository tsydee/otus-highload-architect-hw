-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    birthdate DATE NOT NULL,
    biography TEXT NOT NULL DEFAULT '',
    city VARCHAR(255) NOT NULL,
    gender VARCHAR(16) NOT NULL CHECK (gender IN ('male', 'female')),
    password VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
