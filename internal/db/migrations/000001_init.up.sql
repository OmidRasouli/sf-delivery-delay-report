-- Migration for creating the Vendor table
CREATE TABLE vendors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Migration for creating the User table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


-- Migration for creating the Order table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    delivery_time TIMESTAMPTZ NOT NULL,
    vendor_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Adding foreign key constraint to link orders with users
ALTER TABLE orders ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id);

-- Adding foreign key constraint to link orders with vendors
ALTER TABLE orders ADD CONSTRAINT fk_vendor FOREIGN KEY (vendor_id) REFERENCES vendors (id);


-- Migration for creating the Agent table
CREATE TABLE agents (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Migration for creating the Trip table
CREATE TABLE trips (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE trips ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id);

-- Migration for creating the DelayReport table
CREATE TABLE delay_reports (
    id SERIAL PRIMARY KEY,
    agent_id INT DEFAULT -1,
    order_id INT NOT NULL,
    vendor_id INT NOT NULL,
    delivery_time TIMESTAMPTZ NOT NULL,
    new_delivery_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
	reason VARCHAR(255),
    solved BOOLEAN DEFAULT false
);

ALTER TABLE delay_reports ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id);

--Insert some Vendors
INSERT INTO vendors (name, email) VALUES
('Caspian Grill', 'caspian@example.com'),
('Taj Mahal Restaurant', 'tajmahal@example.com'),
('Golden Dragon', 'goldendragon@example.com'),
('Mediterranean Delight', 'mediterranean@example.com'),
('Sushi Palace', 'sushi@example.com'),
('Pasta Paradise', 'pasta@example.com'),
('Spice Route', 'spiceroute@example.com'),
('Taste of India', 'indian@example.com'),
('Pizza Italia', 'pizza@example.com'),
('Thai Orchid', 'thaiorchid@example.com'),
('Burger Bistro', 'burgerbistro@example.com'),
('Rice & Noodles', 'rice@example.com'),
('Cheese & Co.', 'cheese@example.com'),
('Fusion Flavors', 'fusion@example.com'),
('Smokehouse BBQ', 'bbq@example.com'),
('Café Mosaic', 'cafe@example.com'),
('Mexican Fiesta', 'mexican@example.com'),
('The Grill House', 'grillhouse@example.com'),
('Vegetarian Haven', 'vegetarian@example.com'),
('Savor Street', 'savor@example.com'),
('Seafood Sensations', 'seafood@example.com'),
('Steak & Sizzle', 'steak@example.com'),
('Gourmet Garden', 'gourmet@example.com'),
('Crepe Corner', 'crepe@example.com'),
('Bakery Bliss', 'bakery@example.com'),
('Wok Wonder', 'wok@example.com'),
('Deli Delights', 'deli@example.com'),
('Cinnamon Spice', 'cinnamon@example.com'),
('Juice Junction', 'juice@example.com');


--Insert some Users
INSERT INTO users (name, email, created_at, updated_at, deleted_at) VALUES
('محمد امین', 'mohammad.amin@example.com', NOW(), NOW(), NULL),
('سارا رحمانی', 'sara.rahmani@example.com', NOW(), NOW(), NULL),
('علی نوری', 'ali.noori@example.com', NOW(), NOW(), NULL),
('نازنین رضایی', 'nazanin.rezaei@example.com', NOW(), NOW(), NULL),
('حسین جعفری', 'hossain.jafari@example.com', NOW(), NOW(), NULL),
('نیما محمدی', 'nima.mohammadi@example.com', NOW(), NOW(), NULL),
('فاطمه علیزاده', 'fateme.alizadeh@example.com', NOW(), NOW(), NULL),
('رضا صادقی', 'reza.sadeghi@example.com', NOW(), NOW(), NULL),
('زهرا کریمی', 'zahra.karimi@example.com', NOW(), NOW(), NULL),
('بابک اکبری', 'babak.akbari@example.com', NOW(), NOW(), NULL),
('آتنا مرادی', 'atena.moradi@example.com', NOW(), NOW(), NULL),
('محمدرضا رستمی', 'mohammadreza.rastami@example.com', NOW(), NOW(), NULL),
('نرگس علیزاده', 'nargess.alizadeh@example.com', NOW(), NOW(), NULL),
('امیر حسینی', 'amir.hosseini@example.com', NOW(), NOW(), NULL),
('سارا احمدی', 'sara.ahmadi@example.com', NOW(), NOW(), NULL),
('سجاد محمودی', 'sajjad.mahmoudi@example.com', NOW(), NOW(), NULL),
('آرمین نجفی', 'armin.najafi@example.com', NOW(), NOW(), NULL),
('پریا موسوی', 'parya.mousavi@example.com', NOW(), NOW(), NULL),
('مریم رحیمی', 'maryam.rahimi@example.com', NOW(), NOW(), NULL),
('محمد حسینی', 'mohammad.hosseini@example.com', NOW(), NOW(), NULL),
('رضا اکبری', 'reza.akbari@example.com', NOW(), NOW(), NULL),
('نیما علیزاده', 'nima.alizadeh@example.com', NOW(), NOW(), NULL),
('ساناز میرزایی', 'sanaz.mirzaii@example.com', NOW(), NOW(), NULL),
('علی احمدی', 'ali.ahmadi@example.com', NOW(), NOW(), NULL),
('محمدرضا اکبری', 'mohammadreza.akbari@example.com', NOW(), NOW(), NULL),
('فاطمه علی نژاد', 'fateme.alinezhad@example.com', NOW(), NOW(), NULL),
('سامان محمدی', 'saman.mohammadi@example.com', NOW(), NOW(), NULL),
('شیما حسینی', 'shima.hosseini@example.com', NOW(), NOW(), NULL),
('سعیده علیپور', 'saeideh.alipour@example.com', NOW(), NOW(), NULL);

INSERT INTO agents (name, email) VALUES
('محمد اکبری', 'agent1@example.com'),
('سارا محمدی', 'agent2@example.com'),
('علی رضایی', 'agent3@example.com'),
('نازنین حسینی', 'agent4@example.com'),
('حسین احمدی', 'agent5@example.com'),
('نیما نجفی', 'agent6@example.com'),
('فاطمه علیزاده', 'agent7@example.com'),
('رضا موسوی', 'agent8@example.com'),
('زهرا کاظمی', 'agent9@example.com'),
('بابک محمدی', 'agent10@example.com'),
('آتنا جعفری', 'agent11@example.com'),
('محمدرضا احمدی', 'agent12@example.com'),
('نگار نوروزی', 'agent13@example.com'),
('امیرحسین رحمانی', 'agent14@example.com'),
('ساناز اکبری', 'agent15@example.com'),
('نیما صادقی', 'agent16@example.com'),
('پویا اکبری', 'agent17@example.com'),
('زهرا محمودی', 'agent18@example.com'),
('محمدرضا رضایی', 'agent19@example.com'),
('نازنین صفری', 'agent20@example.com'),
('مریم رضایی', 'agent21@example.com'),
('محمد حسنی', 'agent22@example.com'),
('رضا محمدپور', 'agent23@example.com'),
('نیما مهدی‌پور', 'agent24@example.com'),
('سارا میرزادی', 'agent25@example.com'),
('علی اکبریان', 'agent26@example.com'),
('نگار حسینی', 'agent27@example.com'),
('حسین محمدی', 'agent28@example.com'),
('ندا رضایی', 'agent29@example.com'),
('پویا احمدی', 'agent30@example.com');

-- Generate 1,000,000 rows of random orders
INSERT INTO orders (user_id, delivery_time, vendor_id)
SELECT
    ROUND(RANDOM() * 28) + 1 AS user_id,
    NOW() - INTERVAL '1 month' + (RANDOM() * INTERVAL '1 month') AS delivery_time,
    ROUND(RANDOM() * 28) + 1 AS vendor_id
FROM generate_series(1, 1000000);


-- Generate 100,000 delay reports based on orders
INSERT INTO delay_reports (order_id, vendor_id, delivery_time, new_delivery_time, reason)
SELECT
    o.id AS order_id,
    o.vendor_id,
    o.delivery_time,
    o.delivery_time + INTERVAL '15 minutes' + (RANDOM() * 35)::INTEGER * INTERVAL '1 minute' AS new_delivery_time,
    'vendors delay' AS reason
FROM
    orders o
ORDER BY
    RANDOM()
LIMIT
    100000;
