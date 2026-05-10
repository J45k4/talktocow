-- +migrate Up

-- +migrate StatementBegin

create table if not exists files (
    id serial primary key,
    uploaded_by_user_id integer not null references users(id),
    file_name text not null,
    content_type text not null,
    content_hash text not null,
    size_bytes integer not null,
    created_at timestamp not null default current_timestamp
);

create table if not exists diary_entry_pictures (
    id serial primary key,
    diary_entry_id integer not null references diary_entries(id) on delete cascade,
    file_id integer not null references files(id) on delete cascade,
    created_at timestamp not null default current_timestamp
);

create index if not exists diary_entry_pictures_diary_entry_id_idx on diary_entry_pictures(diary_entry_id);
create index if not exists diary_entry_pictures_file_id_idx on diary_entry_pictures(file_id);
create index if not exists files_content_hash_idx on files(content_hash);
create unique index if not exists diary_entry_pictures_diary_entry_id_file_id_idx on diary_entry_pictures(diary_entry_id, file_id);

-- +migrate StatementEnd
