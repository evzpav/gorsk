CREATE TABLE public.locations (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    name text,
    active boolean DEFAULT TRUE,
    address text,
    company_id INTEGER
);
