CREATE TABLE IF NOT EXISTS devices(
    device_token TEXT
    device_type VARCHAR(15)
    user_id FOREIGN KEY REFERENCES users(id)
)