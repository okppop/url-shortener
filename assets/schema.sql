create table if not exists url (
    id serial,
    original_url varchar not null,
    short_path char(10) not null unique,
    view_times integer not null default 0,
    created_at timestamp not null default current_timestamp,
    expired_at timestamp not null,
    primary key(id)
);