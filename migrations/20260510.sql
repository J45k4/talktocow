-- +migrate Up

-- +migrate StatementBegin

alter table diary_entries
    add column if not exists label text null;

create index if not exists diary_entries_label_idx on diary_entries(label);

-- +migrate StatementEnd
