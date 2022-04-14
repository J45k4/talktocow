-- +migrate Up

-- +migrate StatementBegin

create table courses (
	id serial primary key,
	name varchar(255) not null,
	description text,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table homeworks (
	id serial primary key,
	title varchar(255) not null,
	description text,
	due_date timestamp with time zone not null,
	course_id integer not null,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null,
	foreign key (course_id) references courses(id) on delete cascade
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table homework_submissions (
	id serial primary key,
	homework_id integer not null,
	user_id integer not null,
	submission text not null,
	status int default 0,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null,
	foreign key (homework_id) references homeworks(id) on delete cascade,
	foreign key (user_id) references users(id) on delete cascade
);

-- +migrate StatementEnd

-- +migrate StatementBegin

create table homework_submission_comments (
	id serial primary key,
	homework_submission_id integer not null,
	user_id integer not null,
	comment text not null,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null,
	foreign key (homework_submission_id) references homework_submissions(id) on delete cascade,
	foreign key (user_id) references users(id) on delete cascade
);

-- +migrate StatementEnd