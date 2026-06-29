CREATE TABLE IF NOT EXISTS users
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_credentials
(
    user_id       UUID PRIMARY KEY,
    password_hash TEXT        NOT NULL,
    salt          TEXT        NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_credentials_fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE INDEX user_credentials_idx_user_id ON user_credentials (user_id);

CREATE TABLE IF NOT EXISTS account_statuses
(
    user_id    UUID PRIMARY KEY,
    status     SMALLINT    NOT NULL CHECK (status >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT account_statuses_fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE INDEX account_statuses_idx_user_id ON account_statuses (user_id);

CREATE TABLE IF NOT EXISTS account_status_activities
(
    id            UUID PRIMARY KEY,
    user_id       UUID        NOT NULL,
    activity_type SMALLINT    NOT NULL CHECK (activity_type >= 0),
    occurred_at   TIMESTAMPTZ NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT account_status_activities_fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE INDEX account_status_activities_idx_user_id ON account_status_activities (user_id);

CREATE TABLE IF NOT EXISTS account_locks
(
    account_status_activity_id UUID        NOT NULL,
    reason                     TEXT        NOT NULL,
    created_at                 TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT account_status_activities_fk_account_status_activity_id FOREIGN KEY (account_status_activity_id) REFERENCES account_status_activities (id)
);
CREATE INDEX account_locks_idx_account_status_activity_id ON account_locks (account_status_activity_id);

CREATE TABLE IF NOT EXISTS account_unlocks
(
    account_status_activity_id UUID        NOT NULL,
    reason                     TEXT        NOT NULL,
    created_at                 TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT account_status_activities_fk_account_status_activity_id FOREIGN KEY (account_status_activity_id) REFERENCES account_status_activities (id)
);
CREATE INDEX account_unlocks_idx_account_status_activity_id ON account_unlocks (account_status_activity_id);

CREATE TABLE IF NOT EXISTS user_emails
(
    user_id    UUID PRIMARY KEY,
    email      TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_emails_fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE INDEX user_emails_idx_user_id ON user_emails (user_id);

-- down
DROP TABLE IF EXISTS user_emails;

DROP TABLE IF EXISTS account_unlocks;

DROP TABLE IF EXISTS account_locks;

DROP TABLE IF EXISTS account_status_activities;

DROP TABLE IF EXISTS account_statuses;

DROP TABLE IF EXISTS user_credentials;

DROP TABLE IF EXISTS users;

