-- +migrate Up

-- +migrate StatementBegin

create table if not exists diary_entry_videos (
    id serial primary key,
    diary_entry_id integer not null references diary_entries(id) on delete cascade,
    file_id integer not null references files(id) on delete cascade,
    created_at timestamp not null default current_timestamp
);

create index if not exists diary_entry_videos_diary_entry_id_idx on diary_entry_videos(diary_entry_id);
create index if not exists diary_entry_videos_file_id_idx on diary_entry_videos(file_id);
create unique index if not exists diary_entry_videos_diary_entry_id_file_id_idx on diary_entry_videos(diary_entry_id, file_id);

-- +migrate StatementEnd
