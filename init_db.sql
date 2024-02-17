-- Creating DB and User
CREATE USER sf_user WITH PASSWORD 'OmidRasouli' CREATEDB;
CREATE DATABASE sf_delivery_delay_report WITH OWNER = sf_user;
GRANT ALL PRIVILEGES ON DATABASE sf_delivery_delay_report TO sf_user;