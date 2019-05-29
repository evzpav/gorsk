-- CREATE TABLE roles(
--      id SERIAL PRIMARY KEY,
--      access_level INTEGER NULL DEFAULT NULL,
--      name VARCHAR(100) NULL DEFAULT NULL
-- );

CREATE TABLE public.roles (
    id SERIAL PRIMARY KEY,
    access_level INTEGER,
    name text
);
