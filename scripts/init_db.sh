#!/bin/bash

# Install golang-migrate
curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" | tee /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate

# Run the main application
./sf_delivery_report

# Run database migrations
migrate -source file://internal/db/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down -all
migrate -source file://internal/db/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up

