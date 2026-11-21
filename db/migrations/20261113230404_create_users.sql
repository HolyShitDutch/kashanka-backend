-- +goose Up
create extension if not exists citext;

create table users (
  id uuid not null primary key default uuidv7(),
  display_name bytea not null,
  email citext unique,
  inserted_at timestamptz not null default (now()),
  updated_at timestamptz not null default (now())
);

-- +goose Down
drop table users;
