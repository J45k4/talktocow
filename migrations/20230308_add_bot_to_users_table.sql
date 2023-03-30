-- +migrate Up

-- +migrate StatementBegin

alter table users add column bot boolean not null default false;

-- +migrate StatementEnd