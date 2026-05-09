-- +migrate Up

-- +migrate StatementBegin

create table if not exists diary_entry_pictures (
    id serial primary key,
    diary_entry_id integer not null references diary_entries(id) on delete cascade,
    uploaded_by_user_id integer not null references users(id),
    file_name text not null,
    content_type text not null,
    image_data bytea not null,
    created_at timestamp not null default current_timestamp
);

create index if not exists diary_entry_pictures_diary_entry_id_idx on diary_entry_pictures(diary_entry_id);

-- +migrate StatementEnd
