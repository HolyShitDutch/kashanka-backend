-- +goose Up
create extension if not exists citext;

create table users (
  id uuid not null primary key default uuidv7(),
  email citext not null unique,
  password_hash text,
  inserted_at timestamp not null default (now()),
  updated_at timestamp not null default (now())
);

-- +goose Down
drop table users;
