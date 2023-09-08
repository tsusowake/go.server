# up

CREATE TABLE IF NOT EXISTS `users`
(
    `id`         VARCHAR(26) NOT NULL,
    `created_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS `plans`
(
    `id`         VARCHAR(26)      NOT NULL,
    `name`       VARCHAR(512)     NOT NULL,
    `product_id` VARCHAR(512)     NOT NULL,
    `platform`   TINYINT UNSIGNED NOT NULL,
    `price`      INT UNSIGNED     NOT NULL,
    `created_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS `contracts`
(
    `id`         VARCHAR(26)       NOT NULL,
    `user_id`    VARCHAR(26)       NOT NULL,
    `plan_id`    VARCHAR(26)       NOT NULL,
    `type`       SMALLINT UNSIGNED NOT NULL,
    `started_at` TIMESTAMP         NOT NULL,
    `ended_at`   TIMESTAMP         NOT NULL,
    `created_at` TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`plan_id`) REFERENCES `plans` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS `premium_plan_contracts`
(
    `contract_id`     VARCHAR(26) NOT NULL,
    `subscription_id` VARCHAR(26) NOT NULL,
    `created_at`      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`contract_id`),
    FOREIGN KEY (`contract_id`) REFERENCES `contracts` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS `trial_plan_contracts`
(
    `contract_id` VARCHAR(26) NOT NULL,
    `created_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`contract_id`),
    FOREIGN KEY (`contract_id`) REFERENCES `contracts` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS `contract_activities`
(
    `id`           VARCHAR(26)       NOT NULL,
    `contract_id`  VARCHAR(26)       NOT NULL,
    `activated_at` TIMESTAMP         NOT NULL,
    `type`         SMALLINT UNSIGNED NOT NULL,
    `created_at`   TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`contract_id`) REFERENCES `contracts` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS `contract_intents`
(
    `contract_activity_id` VARCHAR(26)  NOT NULL,
    `price`                INT UNSIGNED NOT NULL,
    `created_at`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`contract_activity_id`),
    FOREIGN KEY (`contract_activity_id`) REFERENCES `contract_activities` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS `contract_cancellations`
(
    `contract_activity_id` VARCHAR(26)       NOT NULL,
    `reason`               SMALLINT UNSIGNED NOT NULL,
    `created_at`           TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`           TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`contract_activity_id`),
    FOREIGN KEY (`contract_activity_id`) REFERENCES `contract_activities` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

CREATE TABLE IF NOT EXISTS `contract_discounts`
(
    `contract_activity_id` VARCHAR(26)       NOT NULL,
    `original_price`       INT UNSIGNED      NOT NULL,
    `discount_amount`      INT UNSIGNED      NOT NULL,
    `reason`               SMALLINT UNSIGNED NOT NULL,
    `gross_price`          INT UNSIGNED      NOT NULL,
    `created_at`           TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`           TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`contract_activity_id`),
    FOREIGN KEY (`contract_activity_id`) REFERENCES `contract_activities` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
;

# down
DROP TABLE IF EXISTS `contract_discounts`;
DROP TABLE IF EXISTS `contract_cancellations`;
DROP TABLE IF EXISTS `contract_intents`;
DROP TABLE IF EXISTS `contract_activities`;
DROP TABLE IF EXISTS `trial_plan_contracts`;
DROP TABLE IF EXISTS `premium_plan_contracts`;
DROP TABLE IF EXISTS `contracts`;
DROP TABLE IF EXISTS `plans`;
DROP TABLE IF EXISTS `users`;
