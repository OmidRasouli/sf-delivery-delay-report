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
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
	reason VARCHAR(255),
    solved BOOLEAN DEFAULT false
);

ALTER TABLE delay_reports ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id);

--Insert some Vendors
INSERT INTO vendors (name, email) VALUES
('شهرزاد', 'shahrzad@example.com'),
('پاساژ اصغری', 'pasaj.asghari@example.com'),
('رستوران مهران', 'mehran.restaurant@example.com'),
('کافه نگین', 'cafe.neghin@example.com'),
('رستوران کوثر', 'koosar.restaurant@example.com');

--Insert some Users
INSERT INTO users (name, email, created_at, updated_at, deleted_at) VALUES
('محمد امین', 'mohammad.amin@example.com', NOW(), NOW(), NULL),
('سارا رضایی', 'sara.rezaei@example.com', NOW(), NOW(), NULL),
('علی حسینی', 'ali.hosseini@example.com', NOW(), NOW(), NULL),
('زهرا محمدی', 'zahra.mohammadi@example.com', NOW(), NOW(), NULL),
('رضا خانی', 'reza.khani@example.com', NOW(), NOW(), NULL),
('فاطمه مرادی', 'fatemeh.moradi@example.com', NOW(), NOW(), NULL),
('حسین جعفری', 'hossein.jafari@example.com', NOW(), NOW(), NULL),
('نازنین احمدی', 'nazanin.ahmadi@example.com', NOW(), NOW(), NULL),
('علی اکبری', 'ali.akbari@example.com', NOW(), NOW(), NULL),
('لیلا رستمی', 'leila.rostami@example.com', NOW(), NOW(), NULL);

INSERT INTO agents (name, email) VALUES
('ایجنت ۱', 'agent1@example.com'),
('ایجنت ۲', 'agent2@example.com'),
('ایجنت ۳', 'agent3@example.com'),
('ایجنت ۴', 'agent4@example.com'),
('ایجنت ۵', 'agent5@example.com');
