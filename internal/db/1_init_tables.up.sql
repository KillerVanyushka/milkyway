CREATE TABLE users
(
    Id          SERIAL   PRIMARY KEY,
    Username    TEXT NOT NULL,
    PhoneNumber TEXT NOT NULL UNIQUE,
    Age         INT,
    Gender      TEXT,
    Email       TEXT NOT NULL UNIQUE,
    Password    TEXT NOT NULL
);