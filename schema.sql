DROP DATABASE IF EXISTS go_htmx_project;

CREATE DATABASE go_htmx_project;

USE go_htmx_project;

CREATE TABLE products(
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(1024) NOT NULL
);