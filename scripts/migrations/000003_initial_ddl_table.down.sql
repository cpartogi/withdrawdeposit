BEGIN;

DROP TABLE IF EXISTS "public"."users" CASCADE;
DROP TABLE IF EXISTS "public"."user_statuses" CASCADE;
DROP TABLE IF EXISTS "public"."user_types" CASCADE;
DROP TABLE IF EXISTS "public"."user_verifications" CASCADE;
DROP TABLE IF EXISTS "public"."user_banks" CASCADE;

COMMIT;