package database

const DDLUsers  =`CREATE TABLE IF NOT EXISTS tg_users
	(
	id        serial               not null,
	username  varchar(100),
	chat_id   integer              not null,
	is_active boolean default true not null,
	user_id   integer              not null
		constraint tg_users_pk
		primary key,
		unique (id),
		unique (chat_id),
		unique (user_id)
	);`

const DDLNotifications  =`
	CREATE TABLE IF NOT EXISTS tg_notifications
	(
	chat_id      integer                                         not null,
	leaving_time timestamp,
	arrival_time timestamp,
	is_sent      boolean default false                           not null,
	created_at   timestamp default now(),
	updated_at   timestamp default now(),
	username     varchar(50)                                     not null,
	msg_id       varchar(16)                                     not null
		constraint tg_notifications_pk
		primary key,
	status       varchar(26) default 'late' :: character varying not null,
	checked_at   timestamp,
	is_deleted   boolean default false                           not null,
		unique (msg_id)
	);`
