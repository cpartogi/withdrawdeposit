BEGIN;

DROP EXTENSION IF EXISTS "uuid-ossp";
DROP FUNCTION IF EXISTS public.id_generator();
DROP SEQUENCE IF EXISTS public.global_id_sequence;

COMMIT;