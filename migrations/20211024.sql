-- +migrate Up

-- +migrate StatementBegin

create table diary_entries (
    id serial primary key,
    title varchar(60),
    body text,
    who_posted_user_id int not null,
    created_at timestamp not null,
    constraint fk_user foreign key(who_posted_user_id) references users(id)
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table shared_diary_entries (
    id serial primary key,
    title varchar(60),
    diary_entry_id int not null,
    user_id int not null,
    created_at timestamp not null,
    constraint fk_user foreign key(user_id) references users(id),
    constraint fk_diary_entry foreign key(diary_entry_id) references diary_entries(id)
)

-- +migrate StatementEnd