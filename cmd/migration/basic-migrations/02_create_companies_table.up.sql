-- CREATE TABLE companies (
--     id SERIAL PRIMARY KEY,
--     created_at  TIMESTAMPTZ DEFAULT NOW(),
-- 	updated_at TIMESTAMPTZ DEFAULT NOW(),
--     name VARCHAR(200) DEFAULT NULL,
--     user_id INTEGER NULL DEFAULT NULL,
-- 	location_id  INTEGER NULL DEFAULT NULL
-- )

CREATE TABLE public.companies (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    name text,
    active boolean
);