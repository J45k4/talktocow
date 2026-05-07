-- +migrate Up

-- +migrate StatementBegin

create table webauthn_users (
    id serial primary key,
    user_id int not null,
    rpid varchar(512) not null,
    handle bytea not null,
    created_at timestamp not null default now(),
    constraint fk_webauthn_users_user foreign key(user_id) references users(id) on delete cascade,
    constraint webauthn_users_rpid_user_id_key unique(rpid, user_id),
    constraint webauthn_users_rpid_handle_key unique(rpid, handle)
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table webauthn_credentials (
    id serial primary key,
    user_id int not null,
    rpid varchar(512) not null,
    credential_id bytea not null,
    credential_json jsonb not null,
    name varchar(100),
    created_at timestamp not null default now(),
    last_used_at timestamp null,
    constraint fk_webauthn_credentials_user foreign key(user_id) references users(id) on delete cascade,
    constraint webauthn_credentials_rpid_credential_id_key unique(rpid, credential_id)
);

create index webauthn_credentials_rpid_user_id_idx on webauthn_credentials(rpid, user_id);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table webauthn_sessions (
    ceremony_id varchar(86) primary key,
    ceremony_kind varchar(32) not null,
    user_id int null,
    rpid varchar(512) not null,
    session_data bytea not null,
    expires_at timestamp with time zone not null,
    created_at timestamp not null default now(),
    constraint fk_webauthn_sessions_user foreign key(user_id) references users(id) on delete cascade
);

create index webauthn_sessions_expires_at_idx on webauthn_sessions(expires_at);

-- +migrate StatementEnd
