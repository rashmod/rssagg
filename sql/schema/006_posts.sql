-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    description TEXT, 
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    published_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE posts;

