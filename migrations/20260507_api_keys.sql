-- +migrate Up

-- +migrate StatementBegin

create table api_keys (
    id serial primary key,
    user_id int not null,
    name varchar(100) not null,
    prefix varchar(16) not null unique,
    token_hash varchar(64) not null unique,
    created_at timestamp not null default now(),
    last_used_at timestamp null,
    revoked_at timestamp null,
    constraint fk_api_keys_user foreign key(user_id) references users(id) on delete cascade
);

create index api_keys_user_id_idx on api_keys(user_id);
create index api_keys_prefix_idx on api_keys(prefix);
create index api_keys_active_idx on api_keys(user_id, revoked_at);

-- +migrate StatementEnd
