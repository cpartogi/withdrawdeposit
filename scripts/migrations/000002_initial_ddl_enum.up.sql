BEGIN;

DROP TYPE IF EXISTS "public"."gender";

CREATE TYPE "public"."gender" AS ENUM ('M', 'F');

COMMIT;