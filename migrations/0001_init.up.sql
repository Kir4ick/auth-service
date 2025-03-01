create table users (
    id serial,
    uuid uuid default gen_random_uuid() primary key,
    login varchar(255) unique,
    email varchar(255) unique,
    hash varchar(255),

    created_at timestamp without time zone default current_timestamp not null,
    updated_at timestamp without time zone default current_timestamp not null,
    deleted_at timestamp without time zone,

    email_verified_at timestamp without time zone
);

create table password_refresh (
    id serial,
    uuid uuid default gen_random_uuid() primary key,
    user_uuid uuid references users(uuid) on delete cascade on update restrict,
    token varchar(255),

    created_at timestamp without time zone default current_timestamp not null
);

create index password_refresh_tokens ON password_refresh(token)
