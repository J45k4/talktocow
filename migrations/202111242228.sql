-- +migrate Up

-- +migrate StatementBegin

create table diary_entry_files (
    id serial primary key,
    diary_entry_id integer not null references diary_entries(id),
    file_hash integer not null,
    created_at timestamp not null default now
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table 

-- +migrate StatementEnd
