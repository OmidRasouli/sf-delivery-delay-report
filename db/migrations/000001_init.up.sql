-- Migration for creating the Order table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    time_delivery TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Migration for creating the DelayReport table
CREATE TABLE delay_reports (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Migration for creating the Trip table
CREATE TABLE trips (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Migration for creating the EmployeeAssignment table
CREATE TABLE employee_assignments (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    employee_id INT NOT NULL,
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Migration for creating the VendorDelayReport table
CREATE TABLE vendor_delay_reports (
    vendor_id INT NOT NULL,
    delay_minutes INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
