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

CREATE TABLE device_data
(
    id serial PRIMARY KEY,
    deviceId int NOT NULL REFERENCES device_iot(id) ON DELETE CASCADE,
    timeStamp timestamp NOT NULL,
    data_type varchar(50) NOT NULL, -- Тип данных (например, температура, влажность)
    value float NOT NULL
);
