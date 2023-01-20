-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
  "id" bigint PRIMARY KEY,
  "name" text,
  "email" text,
  "password" text,
  "reset_password_token" text,
  "reset_password_expired_at" timestamp,
  "created_by" bigint,
  "updated_by" bigint,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT 'now()',
  "deleted_at" timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
