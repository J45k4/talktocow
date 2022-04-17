-- +migrate Up

-- +migrate StatementBegin

create table course_users (
	course_id int not null,
	user_id int not null,
	role int not null,
	primary key (course_id, user_id),
	foreign key (course_id) references courses(id) on delete cascade,
	foreign key (user_id) references users(id) on delete cascade
);

-- +migrate StatementEnd