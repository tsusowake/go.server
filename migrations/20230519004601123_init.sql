# up
CREATE TABLE IF NOT EXISTS transactions
(
    id         VARCHAR(255) NOT NULL,
    created_at DATETIME(3)  NOT NULL DEFAULT '1970-01-01 00:00:00.001',
    updated_at DATETIME(3)  NOT NULL DEFAULT '1970-01-01 00:00:00.001',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS user
(
    id          VARCHAR(255)     NOT NULL,
    password    VARCHAR(255)     NOT NULL,
    email       VARCHAR(320)     NOT NULL UNIQUE,
    status      TINYINT UNSIGNED NOT NULL DEFAULT 0,
    disabled_at DATETIME(3),
    banned_at   DATETIME(3),
    created_at  DATETIME(3)      NOT NULL DEFAULT '1970-01-01 00:00:00.001',
    updated_at  DATETIME(3)      NOT NULL DEFAULT '1970-01-01 00:00:00.001',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS user_setting
(
    user_id    VARCHAR(255) NOT NULL,
    language   VARCHAR(10)  NOT NULL DEFAULT 'ja',
    created_at DATETIME(3)  NOT NULL DEFAULT '1970-01-01 00:00:00.001',
    updated_at DATETIME(3)  NOT NULL DEFAULT '1970-01-01 00:00:00.001',
    PRIMARY KEY (user_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

# down
# DROP TABLE transactions;