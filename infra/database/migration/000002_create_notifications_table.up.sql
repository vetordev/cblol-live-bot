CREATE TABLE IF NOT EXISTS notifications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    schedule TEXT NOT NULL,
    enable INTEGER NOT NULL
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
