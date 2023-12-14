CREATE TABLE IF NOT EXISTS users
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table if not exists user_locks
(
    user_id    bigint primary key,
    lock_type  smallint  not null check (lock_type >= 0),
    locked_at  timestamp not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    constraint user_locks_fk_user_id foreign key (user_id) references users (id)
);

create table if not exists user_credentials
(
    user_id    bigint primary key,
    password   text      not null,
    salt       text      not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    constraint user_credentials_fk_user_id foreign key (user_id) references users (id)
);

create table if not exists user_emails
(
    user_id    bigint primary key,
    email      text      not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    constraint user_emails_fk_user_id foreign key (user_id) references users (id)
);
create index idx_user_emails_email on user_emails (email);

-- down
-- drop table if exists user_emails;
-- drop table if exists user_credentials;
-- drop table if exists user_locks;
-- drop table if exists users;
