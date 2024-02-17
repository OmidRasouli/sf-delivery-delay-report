## SF Delivery Delay Report

This repository contains the source code for the SF Delivery Delay Report application. The application is designed to manage and report delays in delivery services. Below is an overview of the project structure, modules, and API endpoints.

### Project Structure

The project is organized into several main components:

* **app** : Contains the main application logic, including API handlers, services, and repository implementations.
* **db** : Manages the database models and repository interfaces.
* **handler** : Implements HTTP request handlers for different entities (orders, users, vendors, etc.).
* **service** : Provides business logic services, including order management, user services, and more.
* **repository** : Implements the repository interfaces, defining the database operations for each entity.

### Database

The application uses [GORM](https://gorm.io/) as the ORM library to interact with the database. Database models are defined in the `db/models` package.

### Modules

#### 1. Order Module

* **API Endpoints** :
* `POST /api/orders/register`: Create a new order.
* `GET /api/orders/:id`: Retrieve an order by ID.
* `GET /api/orders/user/:id`: Retrieve a list of orders for a specific user.

#### 2. User Module

* **API Endpoints** :
* `GET /api/users/:id`: Retrieve a list of orders for a specific user.

#### 3. Vendor Module

* **API Endpoints** :
* `GET /api/vendors/all`: Retrieve a list of all vendors.
* `GET /api/vendors/:id`: Retrieve a vendor by ID.
* `GET /api/vendors/:id/delayed-orders`: Retrieve delayed orders for a specific vendor.

#### 4. Report Delay Module

* **API Endpoints** :
* `GET /api/report-delay/:id`: Retrieve delay reports for a specific order.

#### 5. Agent Module

* **API Endpoints** :
* `GET /api/agent/assign/:id`: Assign a delay report to an agent.
* `GET /api/agent/done/:agent_id/:order_id`: Mark a delay report as done.

#### 6. Trip Module

* **API Endpoints** :
* `POST /api/trips`: Create a new trip.

## Getting Started

### Prerequisites

- Go installed on your machine
- A PostgreSQL database

### Installation

1. Create a user in PostgreSQL and then create DB and assign that user to the DB and then grant privileges:

   ```bash
   create user sf_agent with password 'SomePassword' CREATEDB;
   create database sf_delivery_delay_report with owner = sf_agent;
   grant all privileges on database sf_delivery_delay_report to sf_agent;
   ```
2. Clone the repository:

   ```bash
   git clone https://github.com/your-username/sf-delivery-delay-report.git
   ```
3. Change into the project directory:

   ```bash
   cd sf-delivery-delay-report
   ```
4. Run this shell file:

   ```bash
   bash ./init_db.sh
   ```
5. Run the application:

   ```bash
   go run main.go
   ```

The API will be available at [http://localhost:8080](http://localhost:8080/).

### Unit tests:

```bash
migrate -source file://internal/db/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down -all
migrate -source file://internal/db/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up
go test -v -cover ./test
```
