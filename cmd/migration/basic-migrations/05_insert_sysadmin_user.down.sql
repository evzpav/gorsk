INSERT INTO public.users (
    id, 
    created_at, 
    updated_at, 
    first_name, 
    last_name, 
    username, 
    password, 
    email, 
    active, 
    role_id, 
    company_id, 
    location_id) 
VALUES (1, now(),now(),'Admin', 'Admin', 'admin', 'admin', 'evzpav@gmail.com', true, 100, 1, 1);