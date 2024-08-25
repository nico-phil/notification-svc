CREATE TABLE IF NOT EXISTS devices(
    device_token TEXT,
    device_type VARCHAR(15),
    user_id INT,
    CONSTRAINT fk_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);