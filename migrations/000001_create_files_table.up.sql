-- +migrate Up
create table files
(
  id         SERIAL primary key,
  type       TEXT  not null,
  version    TEXT  not null,
  content    BYTEA not null,
  hash       TEXT  not null,
  created_at TIMESTAMP default current_timestamp
);