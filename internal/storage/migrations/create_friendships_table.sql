CREATE TABLE friendships
(
    user_id    bigint NOT NULL REFERENCES users (id),
    friend_id  bigint NOT NULL REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, friend_id)
);