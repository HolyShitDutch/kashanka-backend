-- +goose Up
create table webauthn_credentials (
  id uuid not null primary key default uuidv7(),
  user_id uuid not null references users (id) on delete cascade,
  rp_id text not null,
  kid bytea not null,
  credential_json jsonb not null,
  inserted_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  unique (rp_id, kid)
);

-- +goose Down
drop table webauthn_credentials;
