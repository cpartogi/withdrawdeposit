BEGIN;

INSERT INTO public.user_types ("name", created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES
('personal', extract(epoch from now()), NULL, extract(epoch from now()), NULL, NULL, NULL),
('company', extract(epoch from now()), NULL, extract(epoch from now()), NULL, NULL, NULL),
('admin', extract(epoch from now()), NULL, extract(epoch from now()), NULL, NULL, NULL);

INSERT INTO public.user_statuses ("name", created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES
('unverified', extract(epoch from now()), NULL, extract(epoch from now()), NULL, NULL, NULL),
('verified', extract(epoch from now()), NULL, extract(epoch from now()), NULL, NULL, NULL),
('rejected', extract(epoch from now()), NULL, extract(epoch from now()), NULL, NULL, NULL);

COMMIT;