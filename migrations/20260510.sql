-- +migrate Up

-- +migrate StatementBegin

alter table diary_entries
    add column if not exists label text null,
    add column if not exists starts_at timestamp null,
    add column if not exists ends_at timestamp null,
    add column if not exists all_day boolean not null default true;

create index if not exists diary_entries_label_idx on diary_entries(label);
create index if not exists diary_entries_starts_at_idx on diary_entries(starts_at);
create index if not exists diary_entries_ends_at_idx on diary_entries(ends_at);

-- +migrate StatementEnd
