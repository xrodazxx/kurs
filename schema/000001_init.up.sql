CREATE TABLE users
(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE device_iot
(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    type varchar(255) NOT NULL,
    status varchar(255) NOT NULL
);

CREATE TABLE device_data (
    id SERIAL PRIMARY KEY,
    device_id VARCHAR(50) NOT NULL,
    device_name VARCHAR(100),
    timestamp TIMESTAMP DEFAULT NOW(),
    data_type VARCHAR(50),
    value FLOAT8
);
