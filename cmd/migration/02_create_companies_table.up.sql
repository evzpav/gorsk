
CREATE TABLE public.companies (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    name text,
    active BOOLEAN DEFAULT true
);