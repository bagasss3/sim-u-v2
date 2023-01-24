-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tokens (
    id BIGINT NOT NULL,
    email VARCHAR(191) NOT NULL,
    token TEXT NOT NULL,
    expired_at timestamp NOT NULL,
    created_at timestamp NOT NULL DEFAULT 'now()',
    updated_at timestamp NOT NULL DEFAULT 'now()',
    deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tokens;
-- +goose StatementEnd
