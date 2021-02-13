-- +migrate Up

-- +migrate StatementBegin

create table user_received_message (
    id serial primary key,
    user_id int not null,
    message_id int not null,
    received_at timestamp not null,
    read_at timestamp null,
    constraint fk_message foreign key(message_id) references messages(id),
    constraint fk_user foreign key(user_id) references users(id)
);

alter table messages add column reference varchar(128) null;

-- +migrate StatementEnd