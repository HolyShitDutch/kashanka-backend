-- +goose Up
create type user_token_context as enum (
  'access'
);

create table user_tokens (
  id uuid not null primary key default uuidv7(),
  token bytea not null unique,
  context user_token_context not null,
  inserted_at timestamp not null default now()
);

-- +goose Down
drop table user_tokens;
