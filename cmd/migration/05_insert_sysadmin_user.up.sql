INSERT INTO public.users (
    id, 
    first_name, 
    last_name, 
    username, 
    password, 
    email, 
    active, 
    role_id, 
    company_id, 
    location_id) 
VALUES (1, 'Admin', 'Admin', 'admin',crypt('admin', gen_salt('bf', 8)), 'evzpav@gmail.com', true, 100, 1, 1);