BEGIN;

DROP TABLE IF EXISTS "public"."users" CASCADE;
DROP TABLE IF EXISTS "public"."user_statuses" CASCADE;
DROP TABLE IF EXISTS "public"."user_types" CASCADE;
DROP TABLE IF EXISTS "public"."user_verifications" CASCADE;
DROP TABLE IF EXISTS "public"."user_banks" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    "external_user_id" bigint,

    "name" text NOT NULL,
    "email" text NOT NULL, -- encrypted email
    "msisdn" text, -- encrypted msisdn
    "password" text NOT NULL, -- encrypted password

    "user_status_id" integer NOT NULL,
    
    "profile_image" text,
    
    "gender" gender,

    "user_type_id" bigint NOT NULL,

    "is_msisdn_verified" boolean DEFAULT false,
    "is_email_verified" boolean DEFAULT false,

    "subscribe_id" integer,

    "created_at" bigint,
    "created_by" UUID,
    "updated_at" bigint,
    "updated_by" UUID,
    "deleted_at" bigint,
    "deleted_by" UUID
);

CREATE TABLE IF NOT EXISTS "public"."user_statuses" (
    "id" serial NOT NULL PRIMARY KEY,
    "name" text NOT NULL,

    "created_at" bigint,
    "created_by" UUID,
    "updated_at" bigint,
    "updated_by" UUID,
    "deleted_at" bigint,
    "deleted_by" UUID
);

CREATE TABLE IF NOT EXISTS "public"."user_types" (
    "id" serial NOT NULL PRIMARY KEY,
    "name" text NOT NULL,

    "created_at" bigint,
    "created_by" UUID,
    "updated_at" bigint,
    "updated_by" UUID,
    "deleted_at" bigint,
    "deleted_by" UUID
);


CREATE TABLE IF NOT EXISTS "public"."user_verifications" (
    "id" bigserial NOT NULL PRIMARY KEY,
    "user_id" UUID NOT NULL,

    "verification_code" jsonb,  -- otp
    "email_token" jsonb, 
    "reset_password" jsonb,

    "created_at" bigint,
    "created_by" UUID,
    "updated_at" bigint,
    "updated_by" UUID,
    "deleted_at" bigint,
    "deleted_by" UUID
);

CREATE TABLE IF NOT EXISTS "public"."user_banks" (
    "id" bigint NOT NULL PRIMARY KEY DEFAULT id_generator(),
    "bank_id" int NOT NULL,

    "is_default" boolean NOT NULL DEFAULT false,
    "user_id" UUID NOT NULL,

    "bank_account_number" text NOT NULL, -- encrypted bank account number
    "bank_account_name" text NOT NULL,

    "created_at" bigint,
    "created_by" UUID,
    "updated_at" bigint,
    "updated_by" UUID,
    "deleted_at" bigint,
    "deleted_by" UUID
);

COMMIT;