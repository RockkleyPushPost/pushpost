-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    created_at         timestamp with time zone,
    updated_at         timestamp with time zone,
    deleted_at         timestamp with time zone,
    uuid               uuid
        CONSTRAINT users_uuid_unique UNIQUE PRIMARY KEY,

    name               text,
    email              text
        constraint uni_users_email
            unique,
    password           text,
    age                bigint,
    is_email_verified  boolean,
    verification_token text

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
