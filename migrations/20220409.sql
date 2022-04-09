-- +migrate Up

-- +migrate StatementBegin

create table pushover_tokens (
	id serial primary key,
	user_id integer not null,
	token text not null,
	user_token text not null,
	created_at timestamp without time zone not null default now(),
	updated_at timestamp without time zone not null default now()
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table notification_logs (
	id serial primary key,
	user_id integer not null,
	message text not null,
	notification_type int not null,
	created_at timestamp without time zone not null default now()
);

-- +migrate StatementEnd