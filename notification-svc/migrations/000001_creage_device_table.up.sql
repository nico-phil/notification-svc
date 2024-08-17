CREATE TABLE IF NOT EXISTS devices(
    id serial PRIMARY KEY,
    deviceToken text NOT NULL,
    deviceType text NOT NULL
);