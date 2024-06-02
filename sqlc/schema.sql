CREATE TABLE "public"."account_locks" (
    "account_status_activity_id" uuid NOT NULL,
    "reason" text NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX account_locks_idx_account_status_activity_id ON public.account_locks USING btree (account_status_activity_id);

ALTER TABLE ONLY "public"."account_locks" ADD CONSTRAINT "account_status_activities_fk_account_status_activity_id" FOREIGN KEY ("account_status_activity_id") REFERENCES "public"."account_status_activities" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."account_status_activities" (
    "id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "activity_type" smallint NOT NULL CONSTRAINT account_status_activities_activity_type_check CHECK (activity_type >= 0),
    "occurred_at" timestamp with time zone NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE INDEX account_status_activities_idx_user_id ON public.account_status_activities USING btree (user_id);

ALTER TABLE ONLY "public"."account_status_activities" ADD CONSTRAINT "account_status_activities_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."account_statuses" (
    "user_id" uuid NOT NULL,
    "status" smallint NOT NULL CONSTRAINT account_statuses_status_check CHECK (status >= 0),
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id")
);

CREATE INDEX account_statuses_idx_user_id ON public.account_statuses USING btree (user_id);

ALTER TABLE ONLY "public"."account_statuses" ADD CONSTRAINT "account_statuses_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."account_unlocks" (
    "account_status_activity_id" uuid NOT NULL,
    "reason" text NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX account_unlocks_idx_account_status_activity_id ON public.account_unlocks USING btree (account_status_activity_id);

ALTER TABLE ONLY "public"."account_unlocks" ADD CONSTRAINT "account_status_activities_fk_account_status_activity_id" FOREIGN KEY ("account_status_activity_id") REFERENCES "public"."account_status_activities" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."user_credentials" (
    "user_id" uuid NOT NULL,
    "password_hash" text NOT NULL,
    "salt" text NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id")
);

CREATE INDEX user_credentials_idx_user_id ON public.user_credentials USING btree (user_id);

ALTER TABLE ONLY "public"."user_credentials" ADD CONSTRAINT "user_credentials_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."user_emails" (
    "user_id" uuid NOT NULL,
    "email" text NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id")
);

CREATE INDEX user_emails_idx_user_id ON public.user_emails USING btree (user_id);

ALTER TABLE ONLY "public"."user_emails" ADD CONSTRAINT "user_emails_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."users" (
    "id" uuid NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
