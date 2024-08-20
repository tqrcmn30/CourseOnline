-- Table: users
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(25) UNIQUE,
    user_password VARCHAR(85),
    user_email VARCHAR(25) UNIQUE,
    user_phone VARCHAR(15) UNIQUE,
    user_token VARCHAR(255)
);

-- Table: category
CREATE TABLE category (
    cate_id SERIAL PRIMARY KEY,
    cate_name VARCHAR(25) UNIQUE
);

-- Table: courses
CREATE TABLE courses (
    cours_id SERIAL PRIMARY KEY,
    cours_name VARCHAR(55),
    cours_desc VARCHAR(125),
    cours_author VARCHAR(25),
    cours_price DECIMAL(18,2),
    cours_modified TIMESTAMP,
    cours_cate_id INT REFERENCES category(cate_id)
);

-- Table: courses_images
CREATE TABLE courses_images (
    coim_id SERIAL PRIMARY KEY,
    coim_filename VARCHAR(125),
    coim_default VARCHAR(1) CHECK (coim_default IN ('Y', 'N')),
    coim_reme_id INT
);

-- Table: carts
CREATE TABLE carts (
    cart_id SERIAL PRIMARY KEY,
    cart_user_id INT UNIQUE,
    cart_cours_id INT UNIQUE,
    cart_qty INTEGER,
    cart_price DECIMAL(18,2),
    cart_modified TIMESTAMP,
    cart_status VARCHAR(15) CHECK (cart_status IN ('PENDING', 'ORDERED')),
    cart_cart_id INT REFERENCES courses(cours_id)
);

-- Table: order_courses
CREATE TABLE order_courses (
    usco_id SERIAL PRIMARY KEY,
    usco_purchase_no VARCHAR(25) UNIQUE,
    usco_tax DECIMAL(3,2),
    usco_subtotal DECIMAL(18,2),
    usco_patrx_no VARCHAR(55),
    usco_modified TIMESTAMP,
    usco_user_id INT REFERENCES users(user_id)
);

-- Table: order_courses_detail
CREATE TABLE order_courses_detail (
    ucde_id SERIAL PRIMARY KEY,
    ucde_qty INTEGER,
    ucde_price DECIMAL(18,2),
    ucde_total_price DECIMAL(18,2),
    ucde_usco_id INT REFERENCES order_courses(usco_id),
    ucde_cours_id INT REFERENCES courses(cours_id)
);