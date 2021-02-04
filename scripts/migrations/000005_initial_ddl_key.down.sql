BEGIN;

ALTER TABLE "users" DROP CONSTRAINT users_email_key;
ALTER TABLE "users" DROP CONSTRAINT users_msisdn_key;
ALTER TABLE "users" DROP CONSTRAINT users_user_status_id_fkey;
ALTER TABLE "users" DROP CONSTRAINT users_user_type_id_fkey;

ALTER TABLE "user_banks" DROP CONSTRAINT user_banks_user_id_fkey;

ALTER TABLE "user_verifications" DROP CONSTRAINT user_verifications_user_id_fkey;
ALTER TABLE "user_verifications" DROP CONSTRAINT user_verifications_user_id_key;

COMMIT;