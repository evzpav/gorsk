-- CREATE TABLE locations (
--     id SERIAL PRIMARY KEY,
--     created_at  TIMESTAMPTZ DEFAULT NOW(),
-- 	updated_at TIMESTAMPTZ DEFAULT NOW(),
--     name VARCHAR(200) DEFAULT NULL,
--     active BOOLEAN DEFAULT TRUE,
--     address VARCHAR(300) DEFAULT NULL,
-- 	company_id  INTEGER NULL DEFAULT NULL
-- )
CREATE TABLE public.locations (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    name text,
    active boolean,
    address text,
    company_id INTEGER
);
