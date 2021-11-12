-- +migrate Up

-- +migrate StatementBegin

create table diary_entry_comments (
	id serial primary key,
	comment_text text not null,
	diary_entry_id int not null,
	created_at timestamp not null,
	updated_at timestamp null,
	user_id int not null,
	constraint fk_diary_entry_comments_diary_entry foreign key(diary_entry_id) references diary_entries (id),
	constraint fk_diary_entry_comments_user foreign key(user_id) references users (id)
);

-- +migrate StatementEnd