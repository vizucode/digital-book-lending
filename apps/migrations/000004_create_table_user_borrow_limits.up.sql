CREATE TABLE IF NOT EXISTS user_borrow_limits (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    borrow_count INT NOT NULL DEFAULT 0,
    window_start DATETIME NOT NULL,
    window_end DATETIME NOT NULL,
    CONSTRAINT fk_limit_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE KEY uq_user_window (user_id, window_start, window_end)
);