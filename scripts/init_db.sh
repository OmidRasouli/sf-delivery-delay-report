#!/bin/bash

curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/migrate.list
sudo apt list --upgradable
sudo apt update
sudo apt install -y migrate
migrate -source file://db/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up
