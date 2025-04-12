-- +goose Up
-- +goose StatementBegin
CREATE TABLE friendship_requests
(
    uuid           uuid PRIMARY KEY,
    sender_uuid    uuid      NOT NULL,
    recipient_uuid uuid      NOT NULL,
    status         bigint    NOT NULL CHECK (status IN (0, 1, 2)),
    created_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (sender_uuid) REFERENCES users(uuid),
    FOREIGN KEY (recipient_uuid) REFERENCES users(uuid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE friendship_requests;
-- +goose StatementEnd
