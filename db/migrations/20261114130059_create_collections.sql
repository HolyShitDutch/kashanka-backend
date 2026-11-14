-- +goose Up
create type locale as enum (
  'en', 'nl', 'zh-Hant'
);

create type cefr_level as enum (
  'A1', 'A2', 'B1', 'B2', 'C1', 'C2'
);

create table collections (
  id uuid not null primary key default uuidv7(),
  title text not null,
  slug text not null unique,
  parent_id uuid references collections (id) on delete cascade,
  inserted_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

create table cards (
  id uuid not null primary key default uuidv7(),
  question_language locale not null default 'nl',
  answer_languages locale[] not null default '{en,zh-Hant}',
  question text not null,
  answers jsonb not null,
  inserted_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

create table collections_cards (
  id uuid not null primary key default uuidv7(),
  collection_id uuid not null references collections (id) on delete cascade,
  card_id uuid not null references cards (id) on delete cascade,
  position int not null,
  inserted_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  check (position >= 0),
  unique (collection_id, card_id),
  unique (collection_id, position)
);

-- +goose Down
drop table collections_cards, collections, cards;
drop type locale, cefr_level;
