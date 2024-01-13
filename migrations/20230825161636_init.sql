CREATE TABLE IF NOT EXISTS users
(
    id         bigserial primary key,
    created_at timestamp not null default current_timestamp
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

create table if not exists user_access_tokens
(
    id         bigserial primary key,
    user_id    bigint    not null,
    token      text      not null,
    issued_at  timestamp not null,
    expires_at timestamp not null,
    created_at timestamp not null default current_timestamp,
    constraint user_access_tokens_uk_token unique (token),
    constraint user_access_tokens_fk_user_id foreign key (user_id) references users (id)
);

-- down
-- drop table if exists user_access_tokens;
-- drop table if exists user_emails;
-- drop table if exists user_credentials;
-- drop table if exists user_locks;
-- drop table if exists users;
