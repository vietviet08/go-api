CREATE DATABASE IF NOT EXISTS demo_go;

USE demo_go;

CREATE TABLE IF NOT EXISTS album (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(128) NOT NULL,
  artist VARCHAR(255) NOT NULL,
  price DECIMAL(6,2) NOT NULL
);

INSERT INTO album (title, artist, price) VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);