CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE INDEX idx_username ON users(username);

INSERT INTO users (username, password) VALUES
  ('John Doe', 'password123'),
  ('Jane Doe', 'password456');

INSERT INTO users (username, password) VALUES
  ('demo', 'password'),
  ('cook tim', 'password456');