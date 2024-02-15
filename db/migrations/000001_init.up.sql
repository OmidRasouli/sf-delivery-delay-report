-- Migration for creating the Order table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    time_delivery TIMESTAMPTZ NOT NULL,
    status VARCHAR(255) NOT NULL DEFAULT 'PENDING',
    vendor_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE orders ADD CONSTRAINT fk_vendor FOREIGN KEY (vendor_id) REFERENCES vendors (id);

-- Migration for creating the Vendor table
CREATE TABLE vendors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Migration for creating the Agent table
CREATE TABLE agents (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Migration for creating the Trip table
CREATE TABLE trips (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE trips ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id);

-- Migration for creating the DelayReport table
CREATE TABLE delay_reports (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    reason TEXT NOT NULL,
    reporter VARCHAR(255) NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE delay_reports ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id);
