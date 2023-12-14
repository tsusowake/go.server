CREATE TABLE "public"."user_credentials" (
    "user_id" bigint NOT NULL,
    "password" text NOT NULL,
    "salt" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id")
);

ALTER TABLE ONLY "public"."user_credentials" ADD CONSTRAINT "user_credentials_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."user_emails" (
    "user_id" bigint NOT NULL,
    "email" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id")
);

CREATE INDEX idx_user_emails_email ON public.user_emails USING btree (email);

ALTER TABLE ONLY "public"."user_emails" ADD CONSTRAINT "user_emails_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."user_locks" (
    "user_id" bigint NOT NULL,
    "lock_type" smallint NOT NULL CONSTRAINT user_locks_lock_type_check CHECK (lock_type >= 0),
    "locked_at" timestamp NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id")
);

ALTER TABLE ONLY "public"."user_locks" ADD CONSTRAINT "user_locks_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."users" (
    "id" bigserial NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
