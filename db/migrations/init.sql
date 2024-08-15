CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    f_name varchar(255),
    l_name varchar(255),
    phone_number varchar(20) NOT NULL,
    gender int DEFAULT 0,
    age int DEFAULT 0,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users_images (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid,
    image_url varchar(255),
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    constraint id foreign key (user_id) references users (id)
);





-- CREATE TABLE IF NOT EXISTS doctors_images (
--     id primary key uuid DEFAULT uuid_generate_v4 (),
--     doctor_id uuid references doctors (id),
--     image_url varchar(255),
--     created_at timestamp DEFAULT now(),
--     updated_at timestamp DEFAULT now()
-- )

-- CREATE TABLE IF NOT EXISTS doctors (
--     id primary key uuid DEFAULT uuid_generate_v4 (),
--     specialization uuid references categories (id),
--     experience int DEFAULT 0,
--     created_at timestamp DEFAULT now(),
--     updated_at timestamp DEFAULT now()
-- )

-- CREATE TABLE IF NOT EXISTS categories (
--     id primary key uuid DEFAULT uuid_generate_v4 (),
--     name varchar(255),
--     created_at timestamp DEFAULT now(),
--     updated_at timestamp DEFAULT now()
-- )