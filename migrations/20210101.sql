-- +migrate Up

-- +migrate StatementBegin

create table users (
    id serial primary key,
    name varchar(100),
    username varchar(100),
    password_hash varchar(128),
    created_at timestamp not null
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table chatrooms (
    id serial primary key,
    name varchar(80) null,
    created_at timestamp
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table messages (
    id serial primary key,
    message_text text,
    written_at timestamp not null,
    transmited_at timestamp not null,
    server_received_at timestamp not null,
    user_id int not null,
    chatroom_id int not null,
    platform varchar(70),
    created_at timestamp not null,
    constraint fk_user foreign key(user_id) references users(id),
    constraint fk_chatroom foreign key(chatroom_id) references chatrooms(id)
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table events (
    id serial primary key,
    event_text text,
    created_at timestamp not null
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table login_sessions (
    id serial primary key,
    user_id int not null,
    agent varchar(255) not null,
    created_at timestamp not null,
    constraint fk_user foreign key(user_id) references users(id)
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table chatroom_users (
    id serial primary key,
    user_id int not null,
    chatroom_id int not null,
    created_at timestamp,
    constraint fk_user foreign key(user_id) references users(id),
    constraint fk_chatroom foreign key(chatroom_id) references chatrooms(id)
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table chatroom_events (
    id serial primary key,
    chatroom_id int not null,
    message_id int,
    event_type int not null,
    created_at timestamp not null,
    constraint fk_chatroom foreign key(chatroom_id) references chatrooms(id),
    constraint fk_message foreign key(message_id) references messages(id)
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table user_received_chatroom_events (
    id serial primary key,
    user_id int not null,
    chatroom_event_id int not null,
    received_at timestamp not null,
    server_send_at timestamp null,
    created_at timestamp not null,
    constraint fk_user foreign key(user_id) references users(id),
    constraint fk_chatroom_event foreign key(chatroom_event_id) references chatroom_events(id)
);

-- +migrate StatementEnd