CREATE TABLE quotes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    text TEXT,
    author VARCHAR(30),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    likes INT DEFAULT 0,
    dislikes INT DEFAULT 0
);