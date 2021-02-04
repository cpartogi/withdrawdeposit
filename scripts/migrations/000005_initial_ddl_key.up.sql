BEGIN;

ALTER TABLE "users" ADD UNIQUE("email");
ALTER TABLE "users" ADD UNIQUE("msisdn");
ALTER TABLE "users" ADD UNIQUE("external_user_id");
ALTER TABLE "users" ADD FOREIGN KEY ("user_status_id") REFERENCES "user_statuses" ("id");
ALTER TABLE "users" ADD FOREIGN KEY ("user_type_id") REFERENCES "user_types" ("id");

ALTER TABLE "user_verifications" ADD UNIQUE("user_id");
ALTER TABLE "user_verifications" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_banks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

COMMIT;