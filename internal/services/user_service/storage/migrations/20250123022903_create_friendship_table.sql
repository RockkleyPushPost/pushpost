-- +goose Up
-- +goose StatementBegin
CREATE TABLE friendships
(
    uuid                 uuid,
    user_uuid    uuid NOT NULL REFERENCES users (uuid),
    friend_uuid  uuid NOT NULL REFERENCES users (uuid),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_uuid, friend_uuid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE friendships;
-- +goose StatementEnd
