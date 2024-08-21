CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS CATEGORIES
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR(255),
    created_at TIMESTAMP        DEFAULT now(),
    updated_at TIMESTAMP        DEFAULT now()
);

CREATE TABLE IF NOT EXISTS IMAGES
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    image_url  VARCHAR,
    created_at TIMESTAMP        DEFAULT now(),
    updated_at TIMESTAMP        DEFAULT now()
);

CREATE TABLE IF NOT EXISTS USERS
(
    id           uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    f_name       VARCHAR(255),
    l_name       VARCHAR(255),
    phone_number VARCHAR(20) UNIQUE,
    gender       INT              DEFAULT 0,
    age          INT              DEFAULT 0,
    email        VARCHAR(255) UNIQUE,
    created_at   TIMESTAMP        DEFAULT now(),
    updated_at   TIMESTAMP        DEFAULT now()
);

CREATE TABLE IF NOT EXISTS DOCTORS
(
    id           uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    experience   INT              DEFAULT 0,
    f_name       VARCHAR(255),
    l_name       VARCHAR(255),
    description  TEXT,
    phone_number VARCHAR(20) UNIQUE,
    created_at   TIMESTAMP        DEFAULT now(),
    updated_at   TIMESTAMP        DEFAULT now()
);



CREATE TABLE IF NOT EXISTS APPOINTMENT
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    date       DATE,
    time       TIME,
    created_at TIMESTAMP        DEFAULT now(),
    updated_at TIMESTAMP        DEFAULT now()
);

CREATE TABLE IF NOT EXISTS REVIEW
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    rating     FLOAT,
    content    TEXT,
    created_at TIMESTAMP        DEFAULT now(),
    updated_at TIMESTAMP        DEFAULT now()
);

CREATE TABLE IF NOT EXISTS CATEGORY_DOCTOR
(
    id          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_id uuid REFERENCES CATEGORIES (id),
    doctor_id   uuid REFERENCES DOCTORS (id),
    created_at  TIMESTAMP        DEFAULT now(),
    updated_at  TIMESTAMP        DEFAULT now(),
    CONSTRAINT fk_category_id FOREIGN KEY (category_id) REFERENCES CATEGORIES (id),
    CONSTRAINT fk_doctor_id FOREIGN KEY (doctor_id) REFERENCES DOCTORS (id)
);

CREATE TABLE IF NOT EXISTS USERS_IMAGES
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    uuid REFERENCES USERS (id),
    image_id   uuid REFERENCES IMAGES (id),
    created_at TIMESTAMP        DEFAULT now(),
    updated_at TIMESTAMP        DEFAULT now(),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES USERS (id),
    CONSTRAINT fk_image_id FOREIGN KEY (image_id) REFERENCES IMAGES (id)
);

CREATE TABLE IF NOT EXISTS DOCTORS_IMAGES
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    doctor_id  uuid REFERENCES DOCTORS (id),
    image_id   uuid REFERENCES IMAGES (id),
    created_at TIMESTAMP        DEFAULT now(),
    updated_at TIMESTAMP        DEFAULT now(),
    CONSTRAINT fk_doctor_id FOREIGN KEY (doctor_id) REFERENCES DOCTORS (id),
    CONSTRAINT fk_image_id FOREIGN KEY (image_id) REFERENCES IMAGES (id)
);

CREATE TABLE IF NOT EXISTS APPOINTMENT_DOCTOR
(
    id             uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    appointment_id uuid REFERENCES APPOINTMENT (id),
    doctor_id      uuid REFERENCES DOCTORS (id),
    user_id        uuid REFERENCES USERS (id),
    created_at     TIMESTAMP        DEFAULT now(),
    updated_at     TIMESTAMP        DEFAULT now(),
    CONSTRAINT fk_appointment_id FOREIGN KEY (appointment_id) REFERENCES APPOINTMENT (id),
    CONSTRAINT fk_doctor_id FOREIGN KEY (doctor_id) REFERENCES DOCTORS (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES USERS (id)
);

CREATE TABLE IF NOT EXISTS REVIEW_DOCTOR
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    review_id  uuid REFERENCES REVIEW (id),
    doctor_id  uuid REFERENCES DOCTORS (id),
    user_id    uuid REFERENCES USERS (id),
    created_at TIMESTAMP        DEFAULT now(),
    updated_at TIMESTAMP        DEFAULT now(),
    CONSTRAINT fk_review_id FOREIGN KEY (review_id) REFERENCES REVIEW (id),
    CONSTRAINT fk_doctor_id FOREIGN KEY (doctor_id) REFERENCES DOCTORS (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES USERS (id)
);



