-- +goose Up
create extension if not exists citext;

create table users (
  id uuid not null primary key default uuidv7(),
  display_name_encrypted bytea not null,
  email_encrypted bytea,
  email_hash bytea,
  webauthn_id bytea,
  inserted_at timestamptz not null default (now()),
  updated_at timestamptz not null default (now()),
  check (length(webauthn_id) <= 64),
  check ((email_encrypted is not null) = (email_hash is not null))
);

create unique index on users (webauthn_id) where webauthn_id is not null;
create unique index on users (email_hash) where email_hash is not null;

-- +goose Down
drop table users;
